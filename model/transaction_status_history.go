package model

type TransactionStatusHistory struct {
	Id int64 `json:"id"`
	TransactionId int64 `json:"transaction_id"`
	TransactionUuid string `json:"transaction_uuid"`
	Status int `json:"status"`
	CreatedAt int64 `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt int64 `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}
