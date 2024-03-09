package bridgectrl

import (
	"context"
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/bitcoin"
	"github.com/aura-nw/bitcoin-bridge/config"
	"github.com/aura-nw/bitcoin-bridge/database"
	"github.com/aura-nw/bitcoin-bridge/evm"
)

type BridgeControl struct {
	ctx           context.Context
	ctxCancel     context.CancelFunc
	logger        *slog.Logger
	config        *config.BridgeConfig
	db            *database.DB
	bitcoinClient *bitcoin.Client
	evmClient     *evm.Client
}

func checkConfig(config *config.BridgeConfig) {
	if config == nil {
		panic("nil config")
	}
	if config.DBInfo == nil {
		panic("nil db info")
	}
	if config.BitcoinInfo == nil {
		panic("nil bitcoin info")
	}
	if config.EvmInfo == nil {
		panic("nil evm info")
	}
}

func NewBridgeControl(ctx context.Context, logger *slog.Logger, config *config.BridgeConfig) (*BridgeControl, error) {
	checkConfig(config)
	ctx, ctxCancel := context.WithCancel(ctx)
	bc := &BridgeControl{
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

func (bc *BridgeControl) initDB() error {
	db, err := database.NewDB(bc.ctx, bc.logger, bc.config.DBInfo)
	if err != nil {
		return err
	}
	bc.db = db
	return nil
}

func (bc *BridgeControl) initClients() error {
	bitcoinClient, err := bitcoin.NewClient(bc.logger, bc.config.BitcoinInfo)
	if err != nil {
		return err
	}
	bc.bitcoinClient = bitcoinClient

	evmClient, err := evm.NewClient(bc.logger, bc.config.EvmInfo)
	if err != nil {
		return err
	}
	bc.evmClient = evmClient

	return nil
}

func (bc *BridgeControl) Start() {}
func (bc *BridgeControl) Stop()  {}
