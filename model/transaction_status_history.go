package model

import (
	"gorm.io/gorm"
)

type TransactionStatusHistory struct {
	gorm.Model
	Id int64 `json:"id"`
	TransactionId int64 `json:"transaction_id"`
	TransactionUuid string `json:"transaction_uuid"`
	Status int `json:"status"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}
