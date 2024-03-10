package database_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/aura-nw/bitcoin-bridge/config"
	"github.com/aura-nw/bitcoin-bridge/database"
	"github.com/stretchr/testify/require"
)

func TestOpenDB(t *testing.T) {
	dbConfig := config.DBConfigForTest()
	db, err := database.NewDB(context.TODO(), slog.Default(), dbConfig)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })
}
