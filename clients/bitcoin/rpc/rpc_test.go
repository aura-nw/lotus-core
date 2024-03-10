package rpc_test

import (
	"log/slog"
	"testing"

	"github.com/aura-nw/btc-bridge/clients/bitcoin/rpc"
	"github.com/aura-nw/btc-bridge/config"
)

func TestBasic(t *testing.T) {
	authInfo := config.GetBitcoinInfoTestnet()
	rpcClient, err := rpc.NewClient(slog.Default(), authInfo.Host, authInfo.User, authInfo.Password)
	if err != nil {
		t.Fatal(err)
	}
	block, err := rpcClient.GetBlockCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("block count: ", block)
}
