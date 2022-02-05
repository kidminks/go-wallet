package setup

import (
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/helper"
	"taskpot.com/go_wallet/setup"
	"taskpot.com/go_wallet/test/database"
	"taskpot.com/go_wallet/utils"
	"testing"
)

func TestSetup(t *testing.T) {
	utils.ReadParentDatabaseConfig("../test_db_config.json")
	helper.Logger().Info("Database server :- " + config.ParentDbConf.DbServer)
	helper.Logger().Info("Database Username :- " + config.ParentDbConf.DbUser)
	helper.Logger().Info("Database Name :- " + config.ParentDbConf.DbName)
	db := database.GetDbConnection()
	setup.Setup(db)
}


