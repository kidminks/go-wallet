package setup

import (
	"database/sql"
	"io/ioutil"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/helper"
	"taskpot.com/go_wallet/utils"
)

func getDbConnection() *sql.DB {
	dbDriver := config.MySql
	dbUser := config.ParentDbConf.DbUser
	dbPass := config.ParentDbConf.DbPass
	dbName := config.ParentDbConf.DbName
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func closeDb(db *sql.DB) {
	if err := db.Close();err != nil {
		panic(err)
	}
}

func CreateWalletTable() {
	query, err := ioutil.ReadFile(config.CreateTableSqlPath)
	if err != nil {
		panic(err)
	}
	queryList := utils.SplitString(string(query), ";")
	db := getDbConnection()
	defer closeDb(db)
	for _, s := range queryList {
		if len(s) == 0 {
			continue
		}
		helper.Logger().Info("Executing query :- " + s)
		if _, err := db.Exec(s); err != nil {
			panic(err)
		}
	}
}
