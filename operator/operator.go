package operator

import (
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/bitcoin"
	"github.com/aura-nw/bitcoin-bridge/config"
	"github.com/aura-nw/bitcoin-bridge/evm"
)

type Operator struct {
	logger *slog.Logger
	config *config.OperatorConfig

	bitcoinClient *bitcoin.Client
	evmClient     *evm.Client
}

func NewOperator(logger *slog.Logger, config *config.OperatorConfig) (*Operator, error) {
	op := &Operator{
		logger: logger,
		config: config,
	}
	if err := op.initClients(); err != nil {
		return nil, err
	}
	return op, nil
}

func (op *Operator) initClients() error {
	bitcoinClient, err := bitcoin.NewClient(op.logger, op.config.BitcoinInfo)
	if err != nil {
		return err
	}
	op.bitcoinClient = bitcoinClient

	evmClient, err := evm.NewClient(op.logger, op.config.EvmInfo)
	if err != nil {
		return err
	}
	op.evmClient = evmClient

	return nil
}

func (op *Operator) Start() {}
func (op *Operator) Stop()  {}
