package database

import (
	"log/slog"

	"github.com/aura-nw/btc-bridge/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	defaultBatchSize = 100
)

type ChainView interface {
	GetLastSeenHeight() (int64, error)
}

type BitcoinDB interface {
	// View functions
	ChainView

	// For BTC transfer
	StoreBtcDeposits([]types.BtcDeposit) error
	StoreBtcWithdrawals([]types.BtcWithdrawal) error

	// For incriptions transfer
	StoreInscriptionDeposits([]types.InscriptionDeposit) error
	StoreInscriptionWithdrawals([]types.InscriptionWithdrawal) error
}

type bitcoinDBImpl struct {
	logger *slog.Logger
	db     *gorm.DB
}

var _ BitcoinDB = &bitcoinDBImpl{}

// GetLastSeenHeight implements BitcoinDB.
func (b *bitcoinDBImpl) GetLastSeenHeight() (int64, error) {
	// var lastHeight int64
	// err := b.db.Model(&types.BtcDeposit{}).Select("max(height)").Scan(&lastHeight).Error
	// return lastHeight, err
	return 2579303, nil // for testing
}

// StoreBtcDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcDeposits(deposits []types.BtcDeposit) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_hash"}}, DoNothing: true})
	result := deduped.CreateInBatches(&deposits, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(deposits) {
		b.logger.Warn("ignored btc deposit duplicates", "duplicates", len(deposits)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreBtcWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcWithdrawals(withdrawals []types.BtcWithdrawal) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_hash"}}, DoNothing: true})
	result := deduped.CreateInBatches(&withdrawals, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(withdrawals) {
		b.logger.Warn("ignored btc withdrawal duplicates", "duplicates", len(withdrawals)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreInscriptionDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionDeposits(deposits []types.InscriptionDeposit) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_hash"}}, DoNothing: true})
	result := deduped.CreateInBatches(&deposits, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(deposits) {
		b.logger.Warn("ignored inscription deposit duplicates", "duplicates", len(deposits)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreInscriptionWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionWithdrawals(withdrawals []types.InscriptionWithdrawal) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_hash"}}, DoNothing: true})
	result := deduped.CreateInBatches(&withdrawals, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(withdrawals) {
		b.logger.Warn("ignored inscription deposit duplicates", "duplicates", len(withdrawals)-int(result.RowsAffected))
	}

	return result.Error
}

func newBitcoinDBImpl(logger *slog.Logger, db *gorm.DB) BitcoinDB {
	return &bitcoinDBImpl{
		logger: logger,
		db:     db,
	}
}
