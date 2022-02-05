package setup

import (
	"taskpot.com/go_wallet/database"
	"taskpot.com/go_wallet/helper"
)

func databaseSetup() {
	helper.Logger().Info("Setting wallet database")
	database.CreateWalletTable()
	helper.Logger().Info("Wallet database setup complete")
}