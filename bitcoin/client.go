package bitcoin

import (
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/bitcoin/rpc"
	"github.com/aura-nw/bitcoin-bridge/config"
)

// Client defines Bitcoin client.
type Client struct {
	info   *config.BitcoinInfo
	logger *slog.Logger
	rpc    *rpc.Client
}

func NewClient(logger *slog.Logger, info *config.BitcoinInfo) (*Client, error) {
	rpcClient, err := rpc.NewClient(logger, info.Host, info.Password, info.Password)
	if err != nil {
		return nil, err
	}
	return &Client{
		info:   info,
		logger: logger,
		rpc:    rpcClient,
	}, nil
}
