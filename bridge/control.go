package bridge

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/aura-nw/btc-bridge/clients/bitcoin"
	"github.com/aura-nw/btc-bridge/clients/evm"
	"github.com/aura-nw/btc-bridge/config"
	"github.com/aura-nw/btc-bridge/database"
	"github.com/aura-nw/btc-bridge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Control struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	logger    *slog.Logger
	config    *config.BridgeConfig
	db        *database.DB
	btcClient bitcoin.Client
	evmClient *evm.Client
	envelopes map[string]protos.EnvelopeServiceClient
	wg        sync.WaitGroup
}

func NewControl(ctx context.Context, logger *slog.Logger, config *config.BridgeConfig) (*Control, error) {
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

	if err := c.initEnvelopeClient(); err != nil {
		logger.Error("init envelope client failed", "err", err)
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

	return nil
}

func (c *Control) initEnvelopeClient() error {
	c.envelopes = map[string]protos.EnvelopeServiceClient{}
	for _, opInfo := range c.config.Operators {
		conn, err := grpc.Dial(opInfo.GrpcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return err
		}
		c.envelopes[opInfo.GrpcUrl] = protos.NewEnvelopeServiceClient(conn)
	}
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
			lastHeight, err := c.BitcoinDB().GetLastSeenHeight()
			if err != nil {
				c.logger.Error("watchBitcoin: get last seen heigth from database failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}
			c.logger.Info("watchBitcoin: get btc deposits", "height", lastHeight)
			btcDeposits, err := c.btcClient.GetBtcDeposits(lastHeight, c.config.BitcoinMultisig)
			if err != nil {
				c.logger.Error("watchBitcoin: get btc deposits from btc client failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}
			if len(btcDeposits) == 0 {
				c.logger.Info("watchBitcoin: not found btc deposits", "height", lastHeight)
				time.Sleep(1 * time.Second)
				continue
			}
			c.logger.Info("watchBitcoin: found btc deposits", "height", lastHeight, "len", len(btcDeposits))
			if err := c.BitcoinDB().StoreBtcDeposits(btcDeposits); err != nil {
				c.logger.Error("watchBitcoin: store btc deposits failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}

			c.logger.Info("watchBitcoin: requestEnvelopesAndWait", "height", lastHeight)
			if err := c.requestEnvelopesAndWait(lastHeight); err != nil {
				c.logger.Error("watchBitcoin: store btc deposits failed", "err", err)
				time.Sleep(1 * time.Second)
				continue
			}

			// build create invoice tx and send throught eth client
		}
	}
}

func (c *Control) requestEnvelopesAndWait(height int64) error {
	if len(c.envelopes) == 0 {
		return fmt.Errorf("no envelopes")
	}
	var wg sync.WaitGroup
	wg.Add(len(c.envelopes))
	for _, envelope := range c.envelopes {
		ctx := c.ctx
		envelope := envelope
		go func() {
			defer wg.Done()
			resp, err := envelope.VerifyBitcoinDeposits(ctx, &protos.VerifyBitcoinDepositsRequest{
				Height: height,
			})
			if err != nil {
				c.logger.Error("verify bitcoin deposits failed", "height", height, "err", err)
				return
			}
			c.logger.Info("verify bitcoin deposits response", "resp", resp.String())
		}()
	}
	wg.Wait()
	return nil
}

func (c *Control) watchEvm() {
	defer c.wg.Done()
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
