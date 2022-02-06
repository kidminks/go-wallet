package util

import "github.com/kidminks/go_wallet/model"

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
