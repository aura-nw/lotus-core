package database

import (
	"log/slog"

	"gorm.io/gorm"
)

type ChainView interface {
	GetLastSeenHeight() (uint64, error)
}

type BitcoinDB interface {
	// View functions
	ChainView

	// For BTC transfer
	StoreBTCDeposits() error
	StoreBTCWithdrawals() error

	// For incriptions transfer
	StoreTokenDeposits() error
	StoreTokenWithdrawals() error
}

type bitcoinDBImpl struct {
	logger *slog.Logger
	db     *gorm.DB
}

var _ BitcoinDB = &bitcoinDBImpl{}

// GetLastSeenHeight implements BitcoinDB.
func (b *bitcoinDBImpl) GetLastSeenHeight() (uint64, error) {
	panic("unimplemented")
}

// StoreBTCDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBTCDeposits() error {
	panic("unimplemented")
}

// StoreBTCWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreBTCWithdrawals() error {
	panic("unimplemented")
}

// StoreTokenDeposits implements BitcoinDB.
func (b *bitcoinDBImpl) StoreTokenDeposits() error {
	panic("unimplemented")
}

// StoreTokenWithdrawals implements BitcoinDB.
func (b *bitcoinDBImpl) StoreTokenWithdrawals() error {
	panic("unimplemented")
}

func newBitcoinDBImpl(logger *slog.Logger, db *gorm.DB) BitcoinDB {
	return &bitcoinDBImpl{
		logger: logger,
		db:     db,
	}
}
