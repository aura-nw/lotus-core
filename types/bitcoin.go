package types

import (
	"github.com/btcsuite/btcd/btcutil"
	"time"
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
	InvoiceSuccess    InvoiceStatus = "success"
	InvoiceFailed     InvoiceStatus = "failed"
)

type BtcDeposit struct {
	TxId   string `json:"tx_id" gorm:"primary_key:true;not null;index"`
	Height int64  `json:"height" gorm:"index;not null"`
	Memo   string `json:"memo" gorm:"type:text"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction, maybe null
	Receiver       string `json:"receiver"`
	Sender         string `json:"sender" gorm:"not null"`
	MultisigWallet string `json:"multisig_wallet" gorm:"not null"`
	// Amount is btc amount was send
	Amount     btcutil.Amount `json:"amount" gorm:"not null"`
	Idx        uint32         `json:"idx" gorm:"not null;index"`
	UtxoStatus UTXOStatus     `json:"utxo_status" gorm:"not null; default:'unused';index"`
	Status     InvoiceStatus  `json:"status" gorm:"not null; default:'new';index"`
}

func (BtcDeposit) TableName() string {
	return "btc_deposits"
}

type InscriptionDeposit struct {
	Id             int64         `json:"id" gorm:"primary_key:true;not null;autoIncrement"`
	InscriptionId  string        `json:"inscription_id"`
	Status         InvoiceStatus `json:"status" gorm:"not null; default:'new';index"`
	Receiver       string        `json:"receiver"`
	TxHash         string        `json:"tx_hash" gorm:"unique;not null"`
	Height         int64         `json:"height" gorm:"not null"`
	Number         int64         `json:"number"`
	From           string        `json:"from" gorm:"index;not null"`
	To             string        `json:"to" gorm:"not null"`
	Action         string        `json:"action"`
	DateTime       time.Time     `json:"date_time"`
	Token          string        `json:"token"`
	Amount         string        `json:"amount"`
	TokenType      string        `json:"token_type"`
	ContentType    string        `json:"content_type"`
	ContentPreview string        `json:"content_preview"`
	Memo           string        `json:"memo"`
	UpdatedAt      int64         `gorm:"autoUpdateTime"` // Use unix seconds as updating time
	CreatedAt      int64         `gorm:"autoCreateTime"` // Use unix seconds as creating time
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
