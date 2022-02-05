package database

import "taskpot.com/go_wallet/model"

func treatError(err error) {
	if err != nil {
		panic(err)
	}
}

func CreateWalletTable() {
	db := getDbConnection()
	defer closeDb(db)
	treatError(db.AutoMigrate(&model.Wallet{}))
	treatError(db.AutoMigrate(&model.Transaction{}))
	treatError(db.AutoMigrate(&model.TransactionStatusHistory{}))
}
