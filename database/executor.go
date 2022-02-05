package database

import (
	"gorm.io/gorm"
	"taskpot.com/go_wallet/model"
)

func treatError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateWalletTable(db *gorm.DB) {
	treatError(db.AutoMigrate(&model.Wallet{}))
	treatError(db.AutoMigrate(&model.Transaction{}))
	treatError(db.AutoMigrate(&model.TransactionStatusHistory{}))
}
