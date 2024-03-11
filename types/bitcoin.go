package types

type BtcDeposit struct {
	TxHash string `json:"txHash" gorm:"primary_key:true;column:tx_hash;not null"`
	Height int64  `json:"height" gorm:"column:height;not null"`
	Memo   string `json:"memo" gorm:"column:memo"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction, maybe null
	Receiver       string `json:"receiver" gorm:"column:receiver"`
	Sender         string `json:"sender" gorm:"column:sender;not null"`
	MultisigWallet string `json:"multisig_wallet" gorm:"column:multisig_wallet;not null"`
	// Amount is bitcoin amount was send
	Amount string `json:"amount" gorm:"column:amount;not null"`
}

func (BtcDeposit) TableName() string {
	return "btc_deposits"
}

type TokenDeposit struct {
	TxHash    string `json:"txHash" gorm:"primary_key:true;column:tx_hash;not null"`
	Height    int64  `json:"height" gorm:"column:height;not null"`
	Memo      string `json:"memo" gorm:"column:memo"`
	TokenType string `json:"token_type" gorm:"column:token_type"`
}
