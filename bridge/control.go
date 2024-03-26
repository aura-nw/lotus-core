package bridge

import (
	"bytes"
	"context"
	"encoding/hex"
	"log/slog"
	"sync"
	"time"

	"github.com/aura-nw/btc-bridge-core/clients/bitcoin"
	"github.com/aura-nw/btc-bridge-core/clients/evm"
	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/aura-nw/btc-bridge-core/database"
	"github.com/aura-nw/btc-bridge-core/types"
	"github.com/btcsuite/btcd/btcutil"
)

type Control struct {
	ctx       context.Context
	ctxCancel context.CancelFunc

	logger *slog.Logger
	config *config.Config
	db     *database.DB

	btcClient      bitcoin.Client
	evmClient      evm.Client
	lastHeightBtc  int64
	multisigClient bitcoin.MultiSigClient

	wg sync.WaitGroup
}

func NewControl(ctx context.Context, logger *slog.Logger, config *config.Config) (*Control, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	c := &Control{
		logger:    logger,
		config:    config,
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}
	if err := c.initClients(); err != nil {
		logger.Error("init clients failed", "err", err)
		return nil, err
	}

	if err := c.initDB(); err != nil {
		logger.Error("init DB failed", "err", err)
		return nil, err
	}

	return c, nil
}

func (c *Control) initDB() error {
	db, err := database.NewDB(c.ctx, c.logger, c.config.DB)
	if err != nil {
		return err
	}
	c.db = db
	return nil
}

func (c *Control) initClients() error {
	bitcoinClient, err := bitcoin.NewClient(c.logger, c.config.Bitcoin)
	if err != nil {
		return err
	}
	c.btcClient = bitcoinClient

	evmClient, err := evm.NewClient(c.logger, c.config.Evm)
	if err != nil {
		return err
	}
	c.evmClient = evmClient

	multisigClient, err := bitcoin.NewMultiSigClient(c.logger, c.btcClient, c.config.RedeemScript, c.config.PrivateKey)
	if err != nil {
		return err
	}
	c.multisigClient = multisigClient

	return nil
}

func (c *Control) watchBitcoin() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Duration(c.config.Bitcoin.QueryInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("watchBitcoin: context done")
			return
		case <-ticker.C:
			height := c.lastHeightBtc
			c.logger.Info("watchBitcoin: get btc deposits", "height", height)
			btcDeposits, err := c.btcClient.GetBtcDeposits(height, c.config.BitcoinMultisig, c.config.Bitcoin.MinConfirms)
			if err != nil {
				c.logger.Error("watchBitcoin: get btc deposits from btc client failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if len(btcDeposits) == 0 {
				c.logger.Info("watchBitcoin: not found btc deposits", "height", height)
				time.Sleep(1 * time.Second)
				continue
			}
			c.logger.Info("watchBitcoin: found btc deposits", "height", height, "len", len(btcDeposits))
			if err := c.BitcoinDB().StoreBtcDeposits(btcDeposits); err != nil {
				c.logger.Error("watchBitcoin: store btc deposits failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}

			c.lastHeightBtc++

			// build create invoice tx and send throught eth client
		}
	}
}

func (c *Control) watchEvm() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Duration(c.config.Evm.QueryInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("watchEvm: context done")
			return
		case <-ticker.C:
			btcWithdraw, err := c.evmClient.FilterOutgoingInvoice()
			if err != nil {
				c.logger.Error("watchEvm: filter outgoing invoice failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}

			c.logger.Info("watchEvm: found btc withdraw", "len", len(btcWithdraw))

			// store outcome invoice to db
			if err := c.BitcoinDB().StoreBtcWithdraws(btcWithdraw); err != nil {
				c.logger.Error("watchEvm: store btc withdraw failed", "err", err)
				continue
			}
		}
	}

}

func (c *Control) processOutCome() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Duration(c.config.BridgeInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("processOutCome: context done")
			return
		case <-ticker.C:
			go c.generateRawTx()
			go c.broadcastBtcTx()
		}
	}
}

func (c *Control) generateRawTx() {
	btcWithdraws, err := c.BitcoinDB().GetBtcWithdraws()
	if err != nil {
		c.logger.Error("generateRawTx: get btc withdraws failed", "err", err)
		return
	}
	if len(btcWithdraws) == 0 {
		c.logger.Info("generateRawTx: not found btc withdraws")
		return
	}

	withdrawOutput := make([]types.Output, 0, len(btcWithdraws))
	for _, withdraw := range btcWithdraws {
		addr, _ := btcutil.DecodeAddress(withdraw.Address, c.btcClient.GetChainCfg())
		withdrawOutput = append(withdrawOutput, types.Output{
			Address: addr,
			Amount:  withdraw.Amount,
		})
	}

	rawTx, err := c.multisigClient.CreateBTCRawTx(withdrawOutput)
	if err != nil {
		c.logger.Error("generateRawTx: create btc raw tx failed", "err", err)
		return
	}

	var rawTxBytes bytes.Buffer
	if err := rawTx.Serialize(&rawTxBytes); err != nil {
		c.logger.Error("generateRawTx: serialize btc raw tx failed", "err", err)
		return
	}
	rawTxHex := hex.EncodeToString(rawTxBytes.Bytes())
	// TODO: create evm tx and send raw tx hex to contract
	c.logger.Info("generateRawTx: success", "rawTxHex", rawTxHex)

	// update state of outcome invoice
	for i := range btcWithdraws {
		btcWithdraws[i].Status = types.WithdrawProcessing
	}

	if err := c.BitcoinDB().StoreBtcWithdraws(btcWithdraws); err != nil {
		c.logger.Error("generateRawTx: update btc withdraws status failed", "err", err)
		return
	}
}

func (c *Control) broadcastBtcTx() {
	// TODO: wait interface contract event
	//	1. get state of signing tx from contract
	//	2. aggregate sign tx
	//	3. broadcast sign tx
}

func (c *Control) Start() {
	c.wg.Add(3)
	go c.watchBitcoin()
	go c.watchEvm()
	go c.processOutCome()
}
func (c *Control) Stop() {
	c.ctxCancel()
	c.wg.Wait()
}

func (c *Control) BitcoinDB() database.BitcoinDB {
	return c.db.BitcoinDB
}

func (c *Control) EvmDB() database.EvmDB {
	return c.db.EvmDB
}
