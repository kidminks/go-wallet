package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id int64 `json:"id"`
	Uuid string `gorm:"type:varchar(36);default:(UUID())" json:"uuid"`
	PrimaryWalletId int64  `json:"primary_wallet_id"`
	SecondaryWalletId int64 `json:"secondary_wallet_id"`
	PrimaryWalletUuid string `gorm:"type:varchar(36)" json:"primary_wallet_uuid"`
	SecondaryWalletUuid string `gorm:"type:varchar(36)" json:"secondary_wallet_uuid"`
	Amount float64 `json:"amount"`
	Multiplier float64 `json:"multiplier"`
	StartDate int64 `json:"start_date"`
	CompletionDate int64 `json:"completion_date"`
	PaymentDate int64 `json:"payment_date"`
	PaymentStatus int `json:"payment_status"`
	TransactionMetadata datatypes.JSON `json:"transaction_metadata"`
	Comment string `json:"comment"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
