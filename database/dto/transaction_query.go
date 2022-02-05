package dto

import "taskpot.com/go_wallet/model"

type TransactionQuery struct {
	Transaction model.Transaction
	Sort string
	Limit int
	Offset int
}
