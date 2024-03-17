package bitcoin

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aura-nw/btc-bridge/config"
	"github.com/aura-nw/btc-bridge/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
)

// Client defines Bitcoin client.
type Client interface {
	GetBtcDeposits(height int64, filterAddr string, minConfirms int64) ([]types.BtcDeposit, error)
	GetTokenDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error)
}

type clientImpl struct {
	info      config.BitcoinInfo
	logger    *slog.Logger
	rpcClient *rpcclient.Client
}

var _ Client = &clientImpl{}

func NewClient(logger *slog.Logger, info config.BitcoinInfo) (Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         info.Host,
		User:         info.User,
		Pass:         info.Password,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	rpcClient, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		info:      info,
		logger:    logger,
		rpcClient: rpcClient,
	}, nil
}

func (c *clientImpl) GetBtcDeposits(height int64, filterAddr string, minConfirms int64) ([]types.BtcDeposit, error) {
	blockHash, err := c.rpcClient.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	block, err := c.rpcClient.GetBlockVerboseTx(blockHash)
	if err != nil {
		return nil, err
	}
	var results []types.BtcDeposit
	for _, tx := range block.Tx {
		for _, vout := range tx.Vout {
			if filterAddr == vout.ScriptPubKey.Address {
				memo, err := c.getMemo(&tx)
				if err != nil {
					c.logger.Error("get memo failed", "err", err)
				}
				results = append(results, types.BtcDeposit{
					TxHash:         tx.Hash,
					Height:         height,
					Memo:           memo,
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

func (c *clientImpl) GetTokenDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error) {
	var results []types.InscriptionDeposit
	return results, nil
}

// getMemo returns memo for a btc tx, using vout OP_RETURN
func (c *clientImpl) getMemo(tx *btcjson.TxRawResult) (string, error) {
	if tx == nil {
		return "", fmt.Errorf("nil transaction result")
	}
	var opReturns string
	for _, vOut := range tx.Vout {
		if !strings.EqualFold(vOut.ScriptPubKey.Type, "nulldata") {
			continue
		}
		buf, err := hex.DecodeString(vOut.ScriptPubKey.Hex)
		if err != nil {
			c.logger.Error("fail to hex decode scriptPubKey", "err", err)
			continue
		}

		asm, err := txscript.DisasmString(buf)

		if err != nil {
			c.logger.Error("fail to disasm script pubkey", "err", err)
			continue
		}
		opReturnFields := strings.Fields(asm)
		if len(opReturnFields) == 2 {
			// skip "0" field to avoid log noise
			if opReturnFields[1] == "0" {
				continue
			}

			var decoded []byte
			decoded, err = hex.DecodeString(opReturnFields[1])
			if err != nil {
				c.logger.Error("fail to decode OP_RETURN string: %s", "err", err)
				continue
			}
			opReturns += string(decoded)
		}
	}

	return opReturns, nil

}
