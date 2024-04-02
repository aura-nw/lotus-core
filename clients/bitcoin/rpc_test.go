package bitcoin_test

import (
	"log/slog"
	"testing"

	"github.com/aura-nw/lotus-core/clients/bitcoin"
	"github.com/aura-nw/lotus-core/config"
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

func TestClient(t *testing.T) {
	info := config.GetBitcoinInfoTestnet()
	client, err := bitcoin.NewClient(slog.Default(), info)
	require.NoError(t, err)

	deposits, err := client.GetBtcDeposits(2584891, "tb1qffjkknk645jjfd82ed4asmdm8qvgsqj7vstrt3", 3)
	require.NoError(t, err)

	t.Log("deposits: ", deposits)
}
