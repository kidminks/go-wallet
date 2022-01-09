package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/helper"
	"time"
)

var sqlLock = &sync.Mutex{}
var dbConn *sql.DB

func configure(db *sql.DB) {
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)
	db.SetConnMaxLifetime(time.Hour)
}

func createDbConnection() *sql.DB {
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

func getDbConnection() *sql.DB {
	if sqlLock == nil {
		sqlLock.Lock()
		defer sqlLock.Unlock()
		if dbConn == nil {
			helper.Logger().Info("Creating single instance for sql client.")
			dbConn = createDbConnection()
			configure(dbConn)
		}
	}
	return dbConn
}

func ExecuteStmt(query string, args []interface{}) (sql.Result, error) {
	stmt, err := getDbConnection().Prepare(query)
	if err != nil {
		return nil, err
	}
	return stmt.Exec(args)
}

func QueryRow(query string, args []interface{}) *sql.Row {
	return getDbConnection().QueryRow(query, args)
}
