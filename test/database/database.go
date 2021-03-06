package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"github.com/kidminks/go_wallet/config"
	"github.com/kidminks/go_wallet/helper"
	"github.com/kidminks/go_wallet/utils"
)

func getDbConnection() *gorm.DB {
	utils.ReadParentDatabaseConfig("../test_db_config.json")
	dbUser := config.ParentDbConf.DbUser
	dbPass := config.ParentDbConf.DbPass
	dbName := config.ParentDbConf.DbName
	dbServer := config.ParentDbConf.DbServer
	dsn := dbUser+":"+dbPass+"@tcp("+dbServer+")/"+dbName+"?charset=utf8mb4&parseTime=True"
	helper.Logger().Info("Dsn for connecting to db " + dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func closeDb(db *gorm.DB) {
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err := sqlDb.Close();err != nil {
		panic(err)
	}
}

var databaseLock = &sync.Mutex{}
var db *gorm.DB

func GetDbConnection() *gorm.DB {
	if db == nil {
		databaseLock.Lock()
		defer databaseLock.Unlock()
		if db == nil {
			helper.Logger().Info("Creating new database instance")
			db = getDbConnection()
		}
	}
	return db
}
