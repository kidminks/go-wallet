package dto

import "github.com/kidminks/go_wallet/model"

type TransactionQuery struct {
	Transaction model.Transaction
	Sort string
	Limit int
	Offset int
}
