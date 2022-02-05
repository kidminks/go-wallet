package model

import (
	"taskpot.com/go_wallet/model"
	"testing"
)

func getBasicWallet(uuid string) *model.Wallet {
	return &model.Wallet{
		EntityId: 1,
		EntityUuid: uuid,
		EntityType: "USER",
		WalletType: "PERMANENT",
		MaximumBalance: 1000,
		MinimumBalance: 0,
		Status: 1,
		IsActive: true,
		CurrentBalance: 0,
		CreatedBy: "1",
		UpdatedBy: "1",
	}
}

func TestWallet(t *testing.T) {
	wallet := getBasicWallet("test_wallet_creation")
	if wallet.EntityType != "USER" {
		t.Errorf("Error in creating wallet")
	}
}
