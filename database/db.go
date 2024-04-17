package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aura-nw/lotus-core/config"
	"github.com/aura-nw/lotus-core/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type DB struct {
	logger *slog.Logger
	gormDB *gorm.DB

	BitcoinDB BitcoinDB
	EvmDB     EvmDB
}

func NewDB(ctx context.Context, logger *slog.Logger, dbConfig config.DBInfo) (*DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Name, dbConfig.Port)
	gormConfig := gorm.Config{
		// The indexer will explicitly manage the transactions
		SkipDefaultTransaction: true,

		// The postgres parameter counter for a given query is represented with uint16,
		// resulting in a parameter limit of 65535. In order to avoid reaching this limit
		// we'll utilize a batch size of 3k for inserts, well below the limit as long as
		// the number of columns < 20.
		CreateBatchSize: 3_000,

		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	logger.Info("NewDB: opened gorm db")

	// Auto migrations
	if err := autoMigrate(gormDB); err != nil {
		return nil, fmt.Errorf("auto migrate failed: %w", err)
	}

	return &DB{
		logger: logger,
		gormDB: gormDB,

		BitcoinDB: newBitcoinDBImpl(logger, gormDB),
		EvmDB:     newEvmDBImpl(logger, gormDB),
	}, nil
}

func autoMigrate(gormDB *gorm.DB) error {
	return gormDB.AutoMigrate(
		&types.BtcDeposit{},
		&types.BtcWithdraw{},
		&types.InscriptionDeposit{},
		&types.InscriptionWithdrawal{},
	)
}

func (db *DB) Close() error {
	db.logger.Info("closing database")
	sql, err := db.gormDB.DB()
	if err != nil {
		return err
	}

	return sql.Close()
}
