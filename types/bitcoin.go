package types

import (
	"github.com/btcsuite/btcd/btcutil"
)

type UTXOStatus string

const (
	Used   UTXOStatus = "used"
	UnUsed UTXOStatus = "unused"
)

type InvoiceStatus string

const (
	InvoiceNew        InvoiceStatus = "new"
	InvoiceProcessing InvoiceStatus = "processing"
	InvoiceSuccessed  InvoiceStatus = "successed"
	InvoiceFailed     InvoiceStatus = "failed"
)

type BtcDeposit struct {
	TxId   string `json:"tx_id" gorm:"primary_key:true;not null;index"`
	Height int64  `json:"height" gorm:"index;not null"`
	Memo   string `json:"memo"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction, maybe null
	Receiver       string `json:"receiver"`
	Sender         string `json:"sender" gorm:"not null"`
	MultisigWallet string `json:"multisig_wallet" gorm:"not null"`
	// Amount is bitcoin amount was send
	Amount     string        `json:"amount" gorm:"not null"`
	Idx        uint32        `json:"idx" gorm:"not null;index"`
	UtxoStatus UTXOStatus    `json:"utxo_status" gorm:"not null; default:'unused';index"`
	Status     InvoiceStatus `json:"status" gorm:"not null; default:'new';index"`
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
	TxId      string `json:"tx_id" gorm:"primary_key:true;not null;index"`
	Height    int64  `json:"height" gorm:"index;not null"`
	InvoiceId int64  `json:"invoice_id" gorm:"not null"`

	// Receiver is account address of receiver in counterparty chain
	Address string         `json:"address" gorm:"not null"`
	Amount  btcutil.Amount `json:"amount" gorm:"not null"`
	Status  InvoiceStatus  `json:"status" gorm:"not null;default:'new';index"`
}

func (BtcWithdraw) TableName() string {
	return "btc_withdrawals"
}

type InscriptionWithdrawal struct {
}

func (InscriptionWithdrawal) TableName() string {
	return "inscription_withdrawals"
}
