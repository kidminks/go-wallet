package util

import (
	"github.com/kidminks/go_wallet/model"
	"time"
)

func GetBasicTransaction() *model.Transaction {
	return &model.Transaction{
		PrimaryWalletUuid: "primary_wallet_id",
		SecondaryWalletUuid: "secondary_wallet_id",
		Amount: 100,
		Multiplier: 1,
		StartDate: time.Now().Unix(),
		PaymentStatus: 0,
		PaymentDate: time.Now().Unix(),
		Comment: "Test transaction",
	}
}

