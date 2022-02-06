package setup

import (
	"gorm.io/gorm"
	"github.com/kidminks/go_wallet/database"
	"github.com/kidminks/go_wallet/helper"
)

func databaseSetup(db *gorm.DB) {
	helper.Logger().Info("Setting wallet database")
	database.CreateWalletTable(db)
	helper.Logger().Info("Wallet database setup complete")
}