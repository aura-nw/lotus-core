package types

import (
	"github.com/btcsuite/btcd/btcutil"
	"gorm.io/gorm"
	"math/big"
)

type UTXOStatus string

const (
	Used   UTXOStatus = "used"
	UnUsed UTXOStatus = "unused"
)

type WithdrawStatus string

const (
	WithdrawNew        WithdrawStatus = "new"
	WithdrawProcessing WithdrawStatus = "processing"
	WithdrawSuccess    WithdrawStatus = "success"
	WithdrawFailed     WithdrawStatus = "failed"
)

type BtcDeposit struct {
	TxHash string `json:"txHash" gorm:"primary_key:true;not null;index"`
	Height int64  `json:"height" gorm:"index;not null"`
	Memo   string `json:"memo"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction, maybe null
	Receiver       string `json:"receiver"`
	Sender         string `json:"sender" gorm:"not null"`
	MultisigWallet string `json:"multisig_wallet" gorm:"not null"`
	// Amount is bitcoin amount was send
	Amount string     `json:"amount" gorm:"not null"`
	Idx    uint32     `json:"idx" gorm:"not null;index"`
	Status UTXOStatus `json:"status" gorm:"not null; default:'unused';index"`
}

func (BtcDeposit) TableName() string {
	return "btc_deposits"
}

type InscriptionDeposit struct {
	TxHash    string `json:"txHash" gorm:"primary_key:true;not null"`
	Height    int64  `json:"height" gorm:"not null"`
	Memo      string `json:"memo"`
	TokenType string `json:"token_type"`
}

func (InscriptionDeposit) TableName() string {
	return "inscription_deposits"
}

type BtcWithdraw struct {
	gorm.Model

	TxHash    string   `json:"txHash" gorm:"primary_key:true;not null;index"`
	Height    int64    `json:"height" gorm:"index;not null"`
	InvoiceId *big.Int `json:"invoiceId" gorm:"not null"`

	// Receiver is account address of receiver in counterparty chain
	Address string         `json:"address" gorm:"not null"`
	Amount  btcutil.Amount `json:"amount" gorm:"not null"`
	Status  WithdrawStatus `json:"status" gorm:"not null;default:'new';index"`
}

func (BtcWithdraw) TableName() string {
	return "btc_withdrawals"
}

type InscriptionWithdrawal struct {
}

func (InscriptionWithdrawal) TableName() string {
	return "inscription_withdrawals"
}
