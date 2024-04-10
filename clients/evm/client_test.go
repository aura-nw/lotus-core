package evm_test

import (
	"log/slog"
	"testing"

	"github.com/aura-nw/lotus-core/clients/evm"
	"github.com/aura-nw/lotus-core/config"
	"github.com/aura-nw/lotus-core/types"
	"github.com/stretchr/testify/require"
)

func TestCallCreateIncomingInvoice(t *testing.T) {
	testDeposit := types.BtcDeposit{
		TxId:           "12747b5c26bc02d03ab92d9ad8984539b978271941b88e781c772370b5aaf0e123",
		Height:         2574433,
		Memo:           "",
		Receiver:       "0xD02c8cebc86Bd8Cc5fE876b4B793256C0d67a887",
		Sender:         "",
		MultisigWallet: "tb1qrvjce6589p2x9zupd8p0dnkq46s8lsh3rau7v5",
		Amount:         602518,
		Idx:            0,
		UtxoStatus:     "unused",
		Status:         "new",
	}

	client, err := evm.NewClient(slog.Default(), config.GetEvmInfoForTest())
	require.NoError(t, err)

	t.Log("sender", client.GetAddress().Hex())

	err = client.CreateIncomingInvoice(&testDeposit)
	require.NoError(t, err)
}

func TestQuery(t *testing.T) {
	client, err := evm.NewClient(slog.Default(), config.GetEvmInfoForTest())
	require.NoError(t, err)
	t.Log("sender", client.GetAddress().Hex())

	count, err := client.GetOutgoingTxCount()
	require.NoError(t, err)
	t.Log("count", count)
}
