package bridge

import (
	"context"
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/clients/bitcoin"
	"github.com/aura-nw/bitcoin-bridge/clients/evm"
	"github.com/aura-nw/bitcoin-bridge/config"
	"github.com/aura-nw/bitcoin-bridge/database"
)

type Control struct {
	ctx           context.Context
	ctxCancel     context.CancelFunc
	logger        *slog.Logger
	config        *config.BridgeConfig
	db            *database.DB
	bitcoinClient *bitcoin.Client
	evmClient     *evm.Client
}

func NewControl(ctx context.Context, logger *slog.Logger, config *config.BridgeConfig) (*Control, error) {
	ctx, ctxCancel := context.WithCancel(ctx)
	bc := &Control{
		logger:    logger,
		config:    config,
		ctx:       ctx,
		ctxCancel: ctxCancel,
	}

	if err := bc.initDB(); err != nil {
		logger.Error("init DB failed", "err", err)
		return nil, err
	}

	if err := bc.initClients(); err != nil {
		logger.Error("init clients failed", "err", err)
		return nil, err
	}

	return bc, nil
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
	c.bitcoinClient = bitcoinClient

	evmClient, err := evm.NewClient(c.logger, c.config.Evm)
	if err != nil {
		return err
	}
	c.evmClient = evmClient

	return nil
}

func (c *Control) Start() {}
func (c *Control) Stop()  {}
