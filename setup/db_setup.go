package setup

import (
	"gorm.io/gorm"
	"taskpot.com/go_wallet/database"
	"taskpot.com/go_wallet/helper"
)

func databaseSetup(db *gorm.DB) {
	helper.Logger().Info("Setting wallet database")
	database.CreateWalletTable(db)
	helper.Logger().Info("Wallet database setup complete")
}