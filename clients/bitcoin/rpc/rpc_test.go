package rpc_test

import (
	"log/slog"
	"testing"

	"github.com/aura-nw/btc-bridge/clients/bitcoin/rpc"
	"github.com/aura-nw/btc-bridge/config"
	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	authInfo := config.GetBitcoinInfoTestnet()
	rpcClient, err := rpc.NewClient(slog.Default(), authInfo.Host, authInfo.User, authInfo.Password)
	require.NoError(t, err)
	block, err := rpcClient.GetBlockCount()
	require.NoError(t, err)
	t.Log("block count: ", block)

	blockHash, err := rpcClient.GetBlockHash(2579303)
	require.NoError(t, err)
	t.Log("block hash: ", blockHash)
}
