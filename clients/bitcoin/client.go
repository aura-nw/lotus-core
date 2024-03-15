package bitcoin

import (
	"fmt"
	"log/slog"

	"github.com/aura-nw/btc-bridge/clients/bitcoin/rpc"
	"github.com/aura-nw/btc-bridge/config"
	"github.com/aura-nw/btc-bridge/types"
)

// Client defines Bitcoin client.
type Client interface {
	GetBtcDeposits(height int64, filterAddr string) ([]types.BtcDeposit, error)
	GetTokenDeposits(height int64, filterAddr string) ([]types.InscriptionDeposit, error)
}

type clientImpl struct {
	info   config.BitcoinInfo
	logger *slog.Logger
	rpc    *rpc.Client
}

var _ Client = &clientImpl{}

func NewClient(logger *slog.Logger, info config.BitcoinInfo) (Client, error) {
	rpcClient, err := rpc.NewClient(logger, info.Host, info.User, info.Password)
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		info:   info,
		logger: logger,
		rpc:    rpcClient,
	}, nil
}

func (c *clientImpl) GetBtcDeposits(height int64, filterAddr string) ([]types.BtcDeposit, error) {
	blockHash, err := c.rpc.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	block, err := c.rpc.GetBlockVerboseTxs(blockHash)
	if err != nil {
		return nil, err
	}
	var results []types.BtcDeposit
	for _, tx := range block.Tx {
		for _, vout := range tx.Vout {
			if filterAddr == vout.ScriptPubKey.Address {
				results = append(results, types.BtcDeposit{
					TxHash:         tx.Hash,
					Height:         height,
					Memo:           "",
					Receiver:       "",
					Sender:         "",
					MultisigWallet: filterAddr,
					Amount:         fmt.Sprintf("%f", vout.Value),
				})
			}
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no txs found for height: %d, filter addr: %s", height, filterAddr)
	}
	return results, nil
}

func (c *clientImpl) GetTokenDeposits(height int64, filterAddr string) ([]types.InscriptionDeposit, error) {
	var results []types.InscriptionDeposit
	return results, nil
}
