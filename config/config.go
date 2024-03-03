package config

import (
	btcchaincfg "github.com/btcsuite/btcd/chaincfg"
)

type Config struct {
	BitcoinNetwork string `json:"bitcoin_network"`
}

func (c Config) GetBitcoinConfig() *btcchaincfg.Params {
	switch c.BitcoinNetwork {
	case "regtest":
		return &btcchaincfg.RegressionNetParams
	case "testnet3":
		return &btcchaincfg.TestNet3Params
	case "mainnet":
		return &btcchaincfg.MainNetParams
	default:
		return nil
	}
}
