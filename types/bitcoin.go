package types

type BtcDeposit struct {
	TxHash string `json:"txHash" gorm:"primary_key:true;not null"`
	Height int64  `json:"height" gorm:"index;not null"`
	Memo   string `json:"memo"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction, maybe null
	Receiver       string `json:"receiver"`
	Sender         string `json:"sender" gorm:"not null"`
	MultisigWallet string `json:"multisig_wallet" gorm:"not null"`
	// Amount is bitcoin amount was send
	Amount string `json:"amount" gorm:"not null"`
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

type BtcWithdrawal struct {
}

func (BtcWithdrawal) TableName() string {
	return "btc_withdrawals"
}

type InscriptionWithdrawal struct {
}

func (InscriptionWithdrawal) TableName() string {
	return "inscription_withdrawals"
}
