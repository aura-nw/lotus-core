package evm

import (
	"log/slog"

	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/aura-nw/btc-bridge-core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	ChainID() int64
	GetAddress() common.Address

	FilterOutgoingInvoice() ([]types.BtcWithdraw, error)
}

type clientImpl struct {
	client *ethclient.Client
	logger *slog.Logger
	info   config.EvmInfo
}

// GetAddress implements Client.
func (c *clientImpl) GetAddress() common.Address {
	panic("unimplemented")
}

// ChainID implements Client.
func (c *clientImpl) ChainID() int64 {
	return c.info.ChainID
}

// FilterOutgoingInvoice implements Client.
func (c *clientImpl) FilterOutgoingInvoice() ([]types.BtcWithdraw, error) {
	// TODO: implement
	return nil, nil
}

func NewClient(logger *slog.Logger, info config.EvmInfo) (Client, error) {
	client, err := ethclient.Dial(info.Url)
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		logger: logger,
		info:   info,
		client: client,
	}, nil
}

var _ Client = &clientImpl{}
