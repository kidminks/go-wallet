package model

import (
	"gorm.io/gorm"
)

type TransactionStatusHistory struct {
	gorm.Model
	Id int64 `gorm:"primaryKey:autoIncrement" json:"id"`
	TransactionId int64 `gorm:"index;not null" json:"transaction_id"`
	TransactionUuid string `gorm:";not null" json:"transaction_uuid"`
	Status int `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
