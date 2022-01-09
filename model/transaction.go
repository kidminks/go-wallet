package model

type Transaction struct {
	Id int64 `json:"id"`
	Uuid string `json:"uuid"`
	PrimaryWalletId int64  `json:"primary_wallet_id"`
	SecondaryWalletId int64 `json:"secondary_wallet_id"`
	PrimaryWalletUuid string `json:"primary_wallet_uuid"`
	SecondaryWalletUuid string `json:"secondary_wallet_uuid"`
	Amount float64 `json:"amount"`
	Multiplier float64 `json:"multiplier"`
	StartDate int64 `json:"start_date"`
	CompletionDate int64 `json:"completion_date"`
	PaymentDate int64 `json:"payment_date"`
	PaymentStatus int `json:"payment_status"`
	TransactionMetadata StringInterfaceMap `json:"transaction_metadata"`
	Comment string `json:"comment"`
	CreatedAt int64 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
