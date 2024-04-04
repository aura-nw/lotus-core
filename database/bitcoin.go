package database

import (
	"errors"
	"log/slog"

	"github.com/aura-nw/lotus-core/types"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	defaultBatchSize = 100
)

type ChainView interface {
	GetLastSeenHeight() (int64, error)
	GetDepositsByStatus(status []types.InvoiceStatus) ([]*types.BtcDeposit, error)
}

type BitcoinDB interface {
	// View functions
	ChainView

	// For BTC transfer
	StoreBtcDeposits([]types.BtcDeposit) error
	UpdateBtcDeposit(*types.BtcDeposit) error
	// StoreBtcWithdraw stores BTC withdraw in EVM database.
	StoreBtcWithdraw(withdrawal types.BtcWithdraw) error
	// StoreBtcWithdraws stores BTC withdraws in EVM database.
	StoreBtcWithdraws(withdrawals []types.BtcWithdraw) error
	// GetBtcWithdraws returns BTC withdraws from EVM database.
	GetBtcWithdraws() ([]types.BtcWithdraw, error)

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
	return 2574433, nil // for testing
}

// UpdateBtcDeposit implements BitcoinDB.
func (b *bitcoinDBImpl) UpdateBtcDeposit(depsosit *types.BtcDeposit) error {
	return b.db.Save(depsosit).Error
}

// GetDepositsByStatus implements BitcoinDB.
func (b *bitcoinDBImpl) GetDepositsByStatus(status []types.InvoiceStatus) ([]*types.BtcDeposit, error) {
	var results []*types.BtcDeposit
	err := b.db.Where("status IN (?)", status).Find(&results).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*types.BtcDeposit{}, nil
		}
	}
	return results, err
}

// StoreBtcDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcDeposits(deposits []types.BtcDeposit) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_id"}}, DoNothing: true})
	result := deduped.CreateInBatches(&deposits, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(deposits) {
		b.logger.Warn("ignored btc deposit duplicates", "duplicates", len(deposits)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreBtcWithdraws implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcWithdraws(withdrawals []types.BtcWithdraw) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_id"}}, DoNothing: true})
	result := deduped.CreateInBatches(&withdrawals, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(withdrawals) {
		b.logger.Warn("ignored btc withdrawal duplicates", "duplicates", len(withdrawals)-int(result.RowsAffected))
	}

	return result.Error
}

// GetBtcWithdraws implements BitcoinDB.
func (b *bitcoinDBImpl) GetBtcWithdraws() ([]types.BtcWithdraw, error) {
	var withdrawals []types.BtcWithdraw
	result := b.db.Where(types.BtcWithdraw{Status: types.InvoiceNew}).Find(&withdrawals).Limit(defaultBatchSize)
	return withdrawals, result.Error
}

// StoreInscriptionDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionDeposits(deposits []types.InscriptionDeposit) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_id"}}, DoNothing: true})
	result := deduped.CreateInBatches(&deposits, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(deposits) {
		b.logger.Warn("ignored inscription deposit duplicates", "duplicates", len(deposits)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreInscriptionWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreInscriptionWithdrawals(withdrawals []types.InscriptionWithdrawal) error {
	deduped := b.db.Clauses(clause.OnConflict{Columns: []clause.Column{{Name: "tx_id"}}, DoNothing: true})
	result := deduped.CreateInBatches(&withdrawals, defaultBatchSize)
	if result.Error == nil && int(result.RowsAffected) < len(withdrawals) {
		b.logger.Warn("ignored inscription deposit duplicates", "duplicates", len(withdrawals)-int(result.RowsAffected))
	}

	return result.Error
}

// StoreBtcWithdraw implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBtcWithdraw(withdrawal types.BtcWithdraw) error {
	return b.db.Create(&withdrawal).Error
}

func newBitcoinDBImpl(logger *slog.Logger, db *gorm.DB) BitcoinDB {
	return &bitcoinDBImpl{
		logger: logger,
		db:     db,
	}
}
