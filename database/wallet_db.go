package database

import (
	"database/sql"
	"taskpot.com/go_wallet/model"
)

func generateCreateWalletSqlQuery(wallet *model.Wallet) (string, []interface{}){
	sqlStmt := CreateWalletQuery
	var args []interface{}
	sqlStmt += CreateWalletQueryValue
	args = append(args, wallet.EntityId, wallet.EntityUuid, wallet.EntityType, wallet.WalletType, wallet.MinimumBalance,
		wallet.MaximumBalance, wallet.Status, wallet.IsActive, wallet.CurrentBalance, wallet.WalletMetadata,
		wallet.UserMetadata, wallet.CreatedBy, wallet.UpdatedBy)
	return sqlStmt, args
}

func CreateWallet(wallet *model.Wallet) (sql.Result, error) {
	sqlStmt, args := generateCreateWalletSqlQuery(wallet)
	return ExecuteStmt(sqlStmt, args)
}

func GetWalletById(id int64) *sql.Row {
	var args []interface{}
	args = append(args, id)
	return QueryRow(SelectWalletByIdQuery, args)
}