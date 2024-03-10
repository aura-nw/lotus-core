package evm

import (
	"log/slog"

	"github.com/aura-nw/btc-bridge/config"
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
