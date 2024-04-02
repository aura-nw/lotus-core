package types

import "github.com/btcsuite/btcd/btcutil"

type Output struct {
	Address btcutil.Address `json:"address"`
	Amount  btcutil.Amount  `json:"amount"`
}

// UTXO is a class with stateless instances that provides information about an unspent output
type UTXO struct {
	// Transaction ID
	TxId string `json:"tx_id"`
	// Output index
	Vout uint32 `json:"vout"`
	// Address
	Address string `json:"address"`
	// The script included in the output
	ScriptPubKey string `json:"script_pub_key"`
	// Amount is amount string of bitcoin
	Amount string `json:"amount"`
}
