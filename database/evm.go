package database

import (
	"log/slog"

	"gorm.io/gorm"
)

type EvmDB interface {
}

type evmDBImpl struct {
	logger *slog.Logger
	db     *gorm.DB
}

var _ EvmDB = &evmDBImpl{}

func newEvmDBImpl(logger *slog.Logger, db *gorm.DB) EvmDB {
	return &evmDBImpl{
		logger: logger,
		db:     db,
	}
}
