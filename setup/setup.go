package setup

import (
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/helper"
	"taskpot.com/go_wallet/utils"
)

func Setup(dbConfigFile string) {
	helper.Logger().Info("Golang wallet library setup")

	utils.ReadParentDatabaseConfig(dbConfigFile)
	helper.Logger().Info("Database server :- " + config.ParentDbConf.DbServer)
	helper.Logger().Info("Database Username :- " + config.ParentDbConf.DbUser)
	helper.Logger().Info("Database Name :- " + config.ParentDbConf.DbName)

	databaseSetup()
}
