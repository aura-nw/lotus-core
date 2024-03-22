package bitcoin

import (
	"encoding/hex"
	"errors"
	"github.com/aura-nw/btc-bridge-core/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"log/slog"
)

type MultiSigClient interface {
	CreateBTCRawTx(withdrawOutput []types.Output) (*wire.MsgTx, error)
	SignTx(rawTx *wire.MsgTx) ([]byte, error)
	AggregateSigs(sigs [][]byte) ([]byte, error)
}

type multiSigClient struct {
	l            *slog.Logger
	c            Client
	redeemScript []byte
	privateKey   *btcec.PrivateKey
	minFee       btcutil.Amount
}

func NewMultiSigClient(l *slog.Logger, c Client, redeem string, privKey string) (MultiSigClient, error) {
	// redeem script
	redeemScript, err := hex.DecodeString(redeem)
	if err != nil {
		return nil, err
	}

	// private key
	privateKey, err := btcutil.DecodeWIF(privKey)
	if err != nil {
		return nil, err
	}

	return &multiSigClient{
		l:            l,
		c:            c,
		redeemScript: redeemScript,
		privateKey:   privateKey.PrivKey,
	}, nil
}

func (m *multiSigClient) CreateBTCRawTx(withdrawOutput []types.Output) (*wire.MsgTx, error) {
	// get unspent utxo
	unspent, err := m.c.ListUnspent()
	if err != nil {
		m.l.Error("[multiSigClient] ListUnspent", "err", err)
		return nil, err
	}

	// build tx input
	txInputs, refundInfo, err := m.buildTxInput(unspent, m.sumWithdrawOutput(withdrawOutput))
	if err != nil {
		m.l.Error("[multiSigClient] buildTxInput", "err", err)
		return nil, err
	}

	if refundInfo != nil && refundInfo.Amount > 0 {
		withdrawOutput = append(withdrawOutput, *refundInfo)
	}
	txOutputs := m.buildTxOutput(withdrawOutput)

	wireTx, err := m.c.CreateRawTransaction(txInputs, txOutputs)
	if err != nil {
		m.l.Error("[multiSigClient] CreateRawTransaction", "err", err)
		return nil, err
	}

	return wireTx, nil
}

func (m *multiSigClient) SignTx(rawTx *wire.MsgTx) ([]byte, error) {
	return txscript.SignatureScript(rawTx, 0, m.redeemScript, txscript.SigHashAll, m.privateKey, true)
}

func (m *multiSigClient) AggregateSigs(sigs [][]byte) ([]byte, error) {
	signature := txscript.NewScriptBuilder()
	signature.AddOp(txscript.OP_FALSE)
	for _, sig := range sigs {
		signature.AddData(sig)
	}
	signature.AddData(m.redeemScript)
	signatureScript, err := signature.Script()
	if err != nil {
		m.l.Error("[multiSigClient] AggregateSigs", "err", err)
		return nil, err
	}

	return signatureScript, nil
}

func (m *multiSigClient) sumWithdrawOutput(withdrawOutput []types.Output) btcutil.Amount {
	var total btcutil.Amount
	for _, o := range withdrawOutput {
		total += o.Amount
	}
	return total
}

func (m *multiSigClient) buildTxInput(unspentUTXO []btcjson.ListUnspentResult, amount btcutil.Amount) ([]btcjson.TransactionInput, *types.Output, error) {
	// TODO: need to handle Ordinal, BRC-20, RUNE
	var inputs []btcjson.TransactionInput
	var total btcutil.Amount
	var refundInfo types.Output
	for _, u := range unspentUTXO {
		inputValue := btcutil.Amount(u.Amount * 100_000_000)
		inputs = append(inputs, btcjson.TransactionInput{
			Txid: u.TxID,
			Vout: u.Vout,
		})
		total += inputValue
		if total >= amount+m.minFee {
			if total > amount+m.minFee {
				// create new utxo output for this address
				addr, _ := btcutil.DecodeAddress(u.Address, m.c.GetChainCfg())
				refund := total - amount
				refundInfo.Address = addr
				refundInfo.Amount = refund
			}
			break
		}
	}
	if total < amount {
		return nil, nil, errors.New("insufficient funds")
	}

	return inputs, &refundInfo, nil

}

func (m *multiSigClient) buildTxOutput(outputs []types.Output) map[btcutil.Address]btcutil.Amount {
	mapOutputs := make(map[btcutil.Address]btcutil.Amount)
	for _, o := range outputs {
		if _, ok := mapOutputs[o.Address]; ok {
			mapOutputs[o.Address] += o.Amount
			continue
		}
		mapOutputs[o.Address] = o.Amount
	}

	return mapOutputs
}
