package model

import (
	"taskpot.com/go_wallet/model"
	"testing"
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

func TestTransaction(t *testing.T) {
	transaction := GetBasicTransaction()
	if transaction.PrimaryWalletUuid != "primary_wallet_id" || transaction.SecondaryWalletUuid != "secondary_wallet_id" {
		t.Errorf("Error in creating transaction")
	}
}
