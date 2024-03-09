package evm

import (
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/config"
)

type Client struct {
	logger *slog.Logger
	info   *config.EvmInfo
}

func NewClient(logger *slog.Logger, info *config.EvmInfo) (*Client, error) {
	return &Client{logger: logger, info: info}, nil
}
