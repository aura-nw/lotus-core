package database_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aura-nw/lotus-core/config"
	"github.com/aura-nw/lotus-core/database"
	"github.com/aura-nw/lotus-core/types"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/stretchr/testify/require"
)

func TestOpenDB(t *testing.T) {
	dbConfig := config.DBConfigForTest()
	db, err := database.NewDB(context.TODO(), slog.Default(), dbConfig)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })
}

func TestSaveDeposits(t *testing.T) {
	dbConfig := config.DBConfigForTest()
	db, err := database.NewDB(context.TODO(), slog.Default(), dbConfig)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })

	amount, err := btcutil.NewAmount(0.1)
	require.NoError(t, err)

	d1 := types.BtcDeposit{
		TxId:           "0b6517f49f02631b5122366d59e16913a7b9994ef14720e26b78642ebcd23fe8",
		Height:         199,
		Memo:           "memo1",
		Receiver:       "tb1qdrjd6r3rd9vuxxcftxyss9wlept7tvfmafmlnq",
		Sender:         "tb1p9mzeflasu7nlqz009wkyx97hyh0ra74alv4xlyc9ncj52njve2vqwp4q3e",
		MultisigWallet: "tb1p9mzeflasu7nlqz009wkyx97hyh0ra74alv4xlyc9ncj52njve2vqwp4q3e",
		Amount:         amount,
	}

	d2 := types.BtcDeposit{
		TxId:           "5bb1d8b7e7ffe71c1cf89d824e14defe566ba275549f521f4850835b26efae31",
		Height:         200,
		Memo:           "memo2",
		Receiver:       "tb1p9mzeflasu7nlqz009wkyx97hyh0ra74alv4xlyc9ncj52njve2vqwp4q3e",
		Sender:         "tb1qdrjd6r3rd9vuxxcftxyss9wlept7tvfmafmlnq",
		MultisigWallet: "tb1p9mzeflasu7nlqz009wkyx97hyh0ra74alv4xlyc9ncj52njve2vqwp4q3e",
		Amount:         amount,
	}

	deposits := []types.BtcDeposit{d1, d2}

	err = db.BitcoinDB.StoreBtcDeposits(deposits)
	require.NoError(t, err)
}
