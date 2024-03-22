package bridge

import (
	"context"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/aura-nw/btc-bridge-core/clients/bitcoin"
	"github.com/aura-nw/btc-bridge-core/clients/evm"
	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/aura-nw/btc-bridge-core/database"
)

type Control struct {
	ctx           context.Context
	ctxCancel     context.CancelFunc
	logger        *slog.Logger
	config        *config.Config
	db            *database.DB
	btcClient     bitcoin.Client
	evmClient     *evm.Client
	lastHeightBtc atomic.Int64
	wg            sync.WaitGroup
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

	lastHeightBtc, err := c.BitcoinDB().GetLastSeenHeight()
	if err != nil {
		logger.Error("get last seen height from bitcoin db failed", "err", err)
		return nil, err
	}
	c.lastHeightBtc.Store(lastHeightBtc)

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

	return nil
}

func (c *Control) watchBitcoin() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Duration(c.config.Bitcoin.Interval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-c.ctx.Done():
			c.logger.Info("watchBitcoin: context done")
			return
		case <-ticker.C:
			height := c.lastHeightBtc.Load()
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

			c.lastHeightBtc.Add(1)

			// build create invoice tx and send throught eth client
		}
	}
}

func (c *Control) watchEvm() {
	defer c.wg.Done()

	ticker := time.NewTicker(time.Duration(c.config.Evm.Interval) * time.Second)
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
				time.Sleep(1 * time.Second)
				continue
			}
		}
	}

}

func (c *Control) Start() {
	c.wg.Add(2)
	go c.watchBitcoin()
	go c.watchEvm()
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
