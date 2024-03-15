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
	return b.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&deposits, defaultBatchSize).Error
}

// StoreBtcWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcWithdrawals(withdrawals []types.BtcWithdrawal) error {
	return b.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&withdrawals, defaultBatchSize).Error
}

// StoreInscriptionDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionDeposits(deposits []types.InscriptionDeposit) error {
	return b.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&deposits, defaultBatchSize).Error
}

// StoreInscriptionWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionWithdrawals(withdrawals []types.InscriptionWithdrawal) error {
	return b.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).CreateInBatches(&withdrawals, defaultBatchSize).Error
}

func newBitcoinDBImpl(logger *slog.Logger, db *gorm.DB) BitcoinDB {
	return &bitcoinDBImpl{
		logger: logger,
		db:     db,
	}
}
