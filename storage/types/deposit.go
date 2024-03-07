package types

type Deposit struct {
	Id     uint64 `json:"id" gorm:"primary_key:true;column:id;auto_increment;not null"`
	TxHash string `json:"txHash" gorm:"column:tx_hash;not null"`
	Height uint64 `json:"height" gorm:"column:height;not null"`
	// Receiver is account address of receiver in counterparty chain
	// Extract from memo of bitcoin tracsaction
	Receiver       string `json:"receiver" gorm:"column:receiver;not null"`
	Sender         string `json:"sender" gorm:"column:sender;not null"`
	MultisigWallet string `json:"multisigWallet" gorm:"column:multisig_wallet;not null"`
	// Amount is bitcoin amount was send
	Amount uint64 `json:"amount" gorm:"column:amount;not null"`
}

func (Deposit) TableName() string {
	return "deposits"
}
