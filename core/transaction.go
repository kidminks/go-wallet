package core

import (
	"errors"
	"fmt"
	"github.com/kidminks/go_wallet/config"
	"github.com/kidminks/go_wallet/database/dto"
	"github.com/kidminks/go_wallet/helper"
	"github.com/kidminks/go_wallet/model"
	"gorm.io/gorm"
	"time"
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

func validateTransaction(transaction *model.Transaction, db *gorm.DB) error {
	if transaction.PrimaryWalletUuid == "" || transaction.SecondaryWalletUuid == "" {
		return errors.New(config.INVALID_DATA)
	}
	var ids []string
	ids = append(ids, transaction.PrimaryWalletUuid)
	ids = append(ids, transaction.SecondaryWalletUuid)
	getWallets, err := AreWalletsPresent(ids, db)
	if err != nil {
		return err
	}
	for _, wallet := range getWallets {
		if wallet.Uuid == transaction.PrimaryWalletUuid {
			transaction.PrimaryWalletId = wallet.Id
		} else {
			transaction.SecondaryWalletId = wallet.Id
		}
	}
	return nil
}

func CreateTransaction(transaction *model.Transaction, db *gorm.DB) (*model.Transaction, error) {
	helper.Logger().Info("Creating transaction from " + transaction.PrimaryWalletUuid + " to " + transaction.SecondaryWalletUuid)
	err := validateTransaction(transaction, db)
	if err != nil {
		return nil,err
	}
	transaction.StartDate = time.Now().Unix()
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

func ChangeStatus(id string, status int, db *gorm.DB) (*model.Transaction,error) {
	helper.Logger().Info("Searching transaction for " + id)
	transaction := &model.Transaction{}
	result := db.Find(transaction, "id = ? || uuid = ?", id, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_FOUND)
	}
	transactionStatus := &model.TransactionStatusHistory{
		Status: status,
		CreatedBy: transaction.UpdatedBy,
		UpdatedBy: transaction.UpdatedBy,
		TransactionId: transaction.Id,
		TransactionUuid: transaction.Uuid,
	}
	if status == 1 {
		transaction.CompletionDate = time.Now().Unix()
	}
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := SubFromWallet(transaction.SecondaryWalletId, transaction.Amount, tx); err != nil {
			transaction.PaymentStatus = 2
			transaction.Comment = transaction.Comment + "|Failed to debit"
			transactionStatus.Status = 2
		}
		if err := AddToWallet(transaction.PrimaryWalletId, transaction.Amount, tx); err != nil {
			transaction.PaymentStatus = 2
			transaction.Comment = transaction.Comment + "|Failed to credit"
			transactionStatus.Status = 2
		}
		if err := tx.Save(transaction).Error; err != nil {
			return err
		}
		if err := tx.Create(transactionStatus).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
