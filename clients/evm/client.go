package evm

import (
	"github.com/aura-nw/btc-bridge-core/types"
	"log/slog"

	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client struct {
	client *ethclient.Client
	logger *slog.Logger
	info   config.EvmInfo
}

func NewClient(logger *slog.Logger, info config.EvmInfo) (*Client, error) {
	client, err := ethclient.Dial(info.Url)
	if err != nil {
		return nil, err
	}
	return &Client{
		logger: logger,
		info:   info,
		client: client,
	}, nil
}

func (c *Client) ChainID() int64 {
	return c.info.ChainID
}

func (c *Client) FilterOutgoingInvoice() ([]types.BtcWithdraw, error) {
	// TODO: implement
	return nil, nil
}
