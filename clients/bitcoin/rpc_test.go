package bitcoin_test

import (
	"testing"

	"github.com/aura-nw/btc-bridge/config"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	info := config.GetBitcoinInfoTestnet()
	connCfg := &rpcclient.ConnConfig{
		Host:         info.Host,
		User:         info.User,
		Pass:         info.Password,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	rpcClient, err := rpcclient.New(connCfg, nil)
	require.NoError(t, err)
	block, err := rpcClient.GetBlockCount()
	require.NoError(t, err)
	t.Log("block count: ", block)

	blockHash, err := rpcClient.GetBlockHash(2579303)
	require.NoError(t, err)
	t.Log("block hash: ", blockHash)
}
