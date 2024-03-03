package bitcoin

import (
	"log/slog"

	"github.com/aura-nw/bitcoin-bridge/bitcoin/rpc"
	"github.com/aura-nw/bitcoin-bridge/config"
	btcecv2 "github.com/btcsuite/btcd/btcec/v2"
)

// Client defines Bitcoin client.
type Client struct {
	cfg     *config.Config
	logger  *slog.Logger
	rpc     *rpc.Client
	pubKey  string
	privKey *btcecv2.PrivateKey
}

func NewClient(cfg *config.Config, logger *slog.Logger) (*Client, error) {
	return &Client{
		cfg:    cfg,
		logger: logger,
	}, nil
}
