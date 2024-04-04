package bitcoin

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aura-nw/lotus-core/config"
	"github.com/aura-nw/lotus-core/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

const (
	tmpAddress = "0xD02c8cebc86Bd8Cc5fE876b4B793256C0d67a887"
)

// Client defines Bitcoin client.
type Client interface {
	GetBtcDeposits(height int64, filterAddr string, minConfirms int64) ([]types.BtcDeposit, error)
	GetInscriptionDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error)
	ListUnspent() ([]btcjson.ListUnspentResult, error)
	GetChainCfg() *chaincfg.Params
	CreateRawTransaction(inputs []btcjson.TransactionInput, outputs map[btcutil.Address]btcutil.Amount) (*wire.MsgTx, error)
	SendRawTransaction(tx *wire.MsgTx) (*chainhash.Hash, error)
}

type Memo struct {
	Receiver string `json:"receiver"`
}

type clientImpl struct {
	info      config.BitcoinInfo
	logger    *slog.Logger
	rpcClient *rpcclient.Client
	cfgChain  *chaincfg.Params
}

var _ Client = &clientImpl{}

func NewClient(logger *slog.Logger, info config.BitcoinInfo) (Client, error) {
	var cfgChain chaincfg.Params
	if info.Network == "mainnet" {
		cfgChain = chaincfg.MainNetParams
	} else if info.Network == "testnet" || info.Network == "testnet3" {
		cfgChain = chaincfg.TestNet3Params
	} else {
		return nil, fmt.Errorf("unsupport bitcoin network: %s", info.Network)
	}
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
	c := &clientImpl{
		info:      info,
		logger:    logger,
		rpcClient: rpcClient,
		cfgChain:  &cfgChain,
	}

	return c, nil
}

func (c *clientImpl) ListUnspent() ([]btcjson.ListUnspentResult, error) {
	return c.rpcClient.ListUnspent()
}

func (c *clientImpl) GetChainCfg() *chaincfg.Params {
	return c.cfgChain
}

func (c *clientImpl) CreateRawTransaction(inputs []btcjson.TransactionInput,
	outputs map[btcutil.Address]btcutil.Amount) (*wire.MsgTx, error) {
	return c.rpcClient.CreateRawTransaction(inputs, outputs, nil)
}

func (c *clientImpl) SendRawTransaction(tx *wire.MsgTx) (*chainhash.Hash, error) {
	return c.rpcClient.SendRawTransaction(tx, true)
}

func parseMemo(_ string) (Memo, error) {
	return Memo{Receiver: tmpAddress}, nil
}

func (c *clientImpl) getBtcDepositsHelper(height int64, filterAddr string, tx *btcjson.TxRawResult) ([]types.BtcDeposit, error) {
	// sender, err := c.getSender(tx)
	// if err != nil {
	// 	c.logger.Error("get sender from btc tx error", "err", err)
	// 	return nil, err
	// }
	memoStr, err := c.getMemo(tx)
	if err != nil {
		c.logger.Error("get memo from btc tx error", "err", err)
		return nil, err
	}
	memo, err := parseMemo(memoStr)
	if err != nil {
		c.logger.Error("parse memo error", "err", err)
		return nil, err
	}

	outputs, err := c.getOutputs(filterAddr, tx)
	if err != nil {
		c.logger.Error("get outputs from btc tx error", "err", err)
		return nil, err
	}

	var results []types.BtcDeposit
	for _, out := range outputs {
		amount, err := btcutil.NewAmount(out.Value)
		if err != nil {
			c.logger.Error("amount not valid", "value", out.Value)
			continue
		}
		deposit := types.BtcDeposit{
			TxId:           tx.Txid,
			Height:         height,
			Memo:           "",
			Receiver:       memo.Receiver,
			Sender:         "",
			MultisigWallet: filterAddr,
			Amount:         amount,
			Idx:            0,
		}
		results = append(results, deposit)
	}

	return results, nil
}

func (c *clientImpl) getOutputs(filterAddr string, tx *btcjson.TxRawResult) ([]btcjson.Vout, error) {
	var results []btcjson.Vout
	for _, vout := range tx.Vout {
		if vout.ScriptPubKey.Address == filterAddr {
			results = append(results, vout)
		}
	}
	return results, nil
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
	// Check confirmations
	if block.Confirmations < minConfirms {
		return nil, fmt.Errorf("block not enough confirmations, got: %d, min_confirms: %d", block.Confirmations, minConfirms)
	}
	var results []types.BtcDeposit
	for _, tx := range block.Tx {
		deposits, err := c.getBtcDepositsHelper(height, filterAddr, &tx)
		if err != nil {
			c.logger.Error("parse btc deposit tx failed")
			continue
		}
		results = append(results, deposits...)
	}
	return results, nil
}

func (c *clientImpl) GetInscriptionDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error) {
	var results []types.InscriptionDeposit
	return results, nil
}

// getMemo returns memo for a btc tx, using vout OP_RETURN
func (c *clientImpl) getMemo(tx *btcjson.TxRawResult) (string, error) {
	var opReturns string
	for _, vOut := range tx.Vout {
		if !strings.EqualFold(vOut.ScriptPubKey.Type, "nulldata") {
			continue
		}
		buf, err := hex.DecodeString(vOut.ScriptPubKey.Hex)
		if err != nil {
			c.logger.Error("fail to hex decode scriptPubKey")
			continue
		}

		var asm string

		asm, err = txscript.DisasmString(buf)

		if err != nil {
			c.logger.Error("fail to disasm script pubkey")
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
				c.logger.Error("fail to decode OP_RETURN string", "value", opReturnFields[1])
				continue
			}
			opReturns += string(decoded)
		}
	}

	return opReturns, nil
}

// getSender returns sender address for a btc tx, using vin:0
func (c *clientImpl) getSender(tx *btcjson.TxRawResult) (string, error) {
	if len(tx.Vin) == 0 {
		return "", fmt.Errorf("no vin available in tx")
	}

	var vout btcjson.Vout
	txHash, err := chainhash.NewHashFromStr(tx.Vin[0].Txid)
	if err != nil {
		return "", err
	}
	vinTx, err := c.rpcClient.GetRawTransactionVerbose(txHash)
	if err != nil {
		return "", fmt.Errorf("fail to query raw tx")
	}
	vout = vinTx.Vout[tx.Vin[0].Vout]
	addresses := c.getAddressesFromScriptPubKey(vout.ScriptPubKey)
	if len(addresses) == 0 {
		return "", fmt.Errorf("no address available in vout")
	}
	address := addresses[0]
	return address, nil

}
func (c *clientImpl) getAddressesFromScriptPubKey(scriptPubKey btcjson.ScriptPubKeyResult) []string {
	addresses := scriptPubKey.Addresses
	if len(addresses) > 0 {
		return addresses
	}

	if len(scriptPubKey.Hex) == 0 {
		return nil
	}
	buf, err := hex.DecodeString(scriptPubKey.Hex)
	if err != nil {
		c.logger.Error("fail to hex decode script pub key")
		return nil
	}
	_, extractedAddresses, _, err := txscript.ExtractPkScriptAddrs(buf, c.cfgChain)
	if err != nil {
		c.logger.Error("fail to extract addresses from script pub key")
		return nil
	}
	for _, item := range extractedAddresses {
		addresses = append(addresses, item.String())
	}
	return addresses
}
