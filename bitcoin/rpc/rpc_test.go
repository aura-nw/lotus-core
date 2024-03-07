package rpc_test

import (
	"log/slog"
	"testing"

	"github.com/aura-nw/bitcoin-bridge/bitcoin/rpc"
)

func TestBasic(t *testing.T) {
	rpcClient, err := rpc.NewClient("127.0.0.1:18332", "rpcuser", "rpcpass", slog.Default())
	if err != nil {
		t.Fatal(err)
	}
	block, err := rpcClient.GetBlockCount()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("block count: ", block)
}
