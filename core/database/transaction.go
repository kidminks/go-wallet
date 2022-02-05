package database

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/database/dto"
	"taskpot.com/go_wallet/helper"
	"taskpot.com/go_wallet/model"
)

func GetTransactionById(id string, db *gorm.DB) (*model.Transaction, error) {
	helper.Logger().Info("Searching transaction for " + id)
	transaction := &model.Transaction{}
	result := db.Find(transaction, "id = ? || uuid = ?", id, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_FOUND)
	}
	return transaction, nil
}

func CreateTransaction(transaction *model.Transaction, db *gorm.DB) (*model.Transaction, error) {
	helper.Logger().Info("Creating transaction from " + transaction.PrimaryWalletUuid + " to " + transaction.SecondaryWalletUuid)
	result := db.Create(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_INSERTED)
	}
	return transaction, nil
}

func FindTransaction(query dto.TransactionQuery, db *gorm.DB) ([]*model.Transaction, error) {
	fmt.Printf("Fetching transation using data %+v\n",query)
	var transactions []*model.Transaction
	result := db.Where(query.Transaction).Order(query.Sort).Limit(query.Limit).Offset(query.Offset).Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
