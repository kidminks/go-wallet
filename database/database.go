package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/helper"
)

func getDbConnection() *gorm.DB {
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
