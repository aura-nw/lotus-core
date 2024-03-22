package types

import "github.com/btcsuite/btcd/btcutil"

type Output struct {
	Address btcutil.Address `json:"address"`
	Amount  btcutil.Amount  `json:"amount"`
}
