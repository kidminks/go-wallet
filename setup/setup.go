package setup

import (
	"taskpot.com/go_payments/config"
	"taskpot.com/go_payments/helper"
	"taskpot.com/go_payments/utils"
)

func Setup(dbConfigFile string) {
	helper.Logger().Info("Golang wallet library setup")

	utils.ReadParentDatabaseConfig(dbConfigFile)
	helper.Logger().Info("Database Username :- " + config.ParentDbConf.DbUser)
	helper.Logger().Info("Database Name :- " + config.ParentDbConf.DbName)

	helper.DatabaseSetup()
}
