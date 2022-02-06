package main

import (
	"github.com/kidminks/go_wallet/config"
	"github.com/kidminks/go_wallet/helper"
	"github.com/kidminks/go_wallet/setup"
	"github.com/kidminks/go_wallet/test/database"
	"github.com/kidminks/go_wallet/utils"
)

func main() {
	utils.ReadParentDatabaseConfig("test/test_db_config.json")
	helper.Logger().Info("Database server :- " + config.ParentDbConf.DbServer)
	helper.Logger().Info("Database Username :- " + config.ParentDbConf.DbUser)
	helper.Logger().Info("Database Name :- " + config.ParentDbConf.DbName)
	db := database.GetDbConnection()
	setup.Setup(db)
}


