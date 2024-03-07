package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/aura-nw/bitcoin-bridge/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	logger *slog.Logger
	gormDB *gorm.DB
}

func NewDB(ctx context.Context, logger *slog.Logger, dbConfig config.DBConfig) (*DB, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Name)
	if dbConfig.Port != 0 {
		dsn += fmt.Sprintf(" port=%d", dbConfig.Port)
	}
	if dbConfig.User != "" {
		dsn += fmt.Sprintf(" user=%s", dbConfig.User)
	}
	if dbConfig.Password != "" {
		dsn += fmt.Sprintf(" password=%s", dbConfig.Password)
	}

	gormConfig := gorm.Config{
		// The indexer will explicitly manage the transactions
		SkipDefaultTransaction: true,

		// The postgres parameter counter for a given query is represented with uint16,
		// resulting in a parameter limit of 65535. In order to avoid reaching this limit
		// we'll utilize a batch size of 3k for inserts, well below the limit as long as
		// the number of columns < 20.
		CreateBatchSize: 3_000,
	}

	gormDB, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &DB{
		logger: logger,
		gormDB: gormDB,
	}, nil
}

func (db *DB) ExecuteSQLMigration(migrationsFolder string) error {
	err := filepath.Walk(migrationsFolder, func(path string, info os.FileInfo, err error) error {
		// Check for any walking error
		if err != nil {
			return fmt.Errorf("failed to process migration file: %s", path)
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Read the migration file content
		db.logger.Info("reading sql file", "path", path)
		fileContent, readErr := os.ReadFile(path)
		if readErr != nil {
			return fmt.Errorf("error reading SQL file: %s", path)
		}

		// Execute the migration
		db.logger.Info("executing sql file", "path", path)
		execErr := db.gormDB.Exec(string(fileContent)).Error
		if execErr != nil {
			return fmt.Errorf("error executing SQL script: %s", path)
		}

		return nil
	})

	db.logger.Info("finished migrations")
	return err
}

func (db *DB) Close() error {
	db.logger.Info("closing database")
	sql, err := db.gormDB.DB()
	if err != nil {
		return err
	}

	return sql.Close()
}