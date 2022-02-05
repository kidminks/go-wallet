package model

import (
	"taskpot.com/go_wallet/model"
	"testing"
)

func GetBasicWallet(uuid string) *model.Wallet {
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

func GetWithWalletUuid(uuid string,userUuid string) *model.Wallet {
	wallet := GetBasicWallet(userUuid)
	wallet.Uuid = uuid
	return wallet
}

func TestWallet(t *testing.T) {
	wallet := GetBasicWallet("test_wallet_creation")
	if wallet.EntityType != "USER" {
		t.Errorf("Error in creating wallet")
	}
}
