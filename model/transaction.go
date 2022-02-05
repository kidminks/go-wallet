package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id int64 `gorm:"primaryKey:autoIncrement" json:"id"`
	Uuid string `gorm:"type:varchar(36);default:(UUID());uniqueIndex" json:"uuid"`
	PrimaryWalletId int64  `gorm:"index;not null" json:"primary_wallet_id"`
	SecondaryWalletId int64 `gorm:"index;not null" json:"secondary_wallet_id"`
	PrimaryWalletUuid string `gorm:"type:varchar(36);index;not null" json:"primary_wallet_uuid"`
	SecondaryWalletUuid string `gorm:"type:varchar(36);index;not null" json:"secondary_wallet_uuid"`
	Amount float64 `gorm:"default:0" json:"amount"`
	Multiplier float64 `gorm:"default:1;not null" json:"multiplier"`
	StartDate int64 `json:"start_date"`
	CompletionDate int64 `json:"completion_date"`
	PaymentDate int64 `gorm:"index;not null" json:"payment_date"`
	PaymentStatus int `gorm:"default:1;not null" json:"payment_status"`
	TransactionMetadata datatypes.JSON `json:"transaction_metadata"`
	Comment string `json:"comment"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
