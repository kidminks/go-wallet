package helper

import "taskpot.com/go_wallet/setup"

func DatabaseSetup() {
	Logger().Info("Setting wallet database")
	setup.CreateWalletTable()
	Logger().Info("Wallet database setup complete")
}