package helper

func DatabaseSetup() {
	Logger().Info("Setting wallet database")
	CreateWalletTable()
	Logger().Info("Wallet database setup complete")
}