package core

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"github.com/kidminks/go_wallet/config"
	"github.com/kidminks/go_wallet/database/dto"
	"github.com/kidminks/go_wallet/helper"
	"github.com/kidminks/go_wallet/model"
)

func GetWalletById(id string, db *gorm.DB) (*model.Wallet, error) {
	helper.Logger().Info("Fetching wallet for id" + id)
	wallet := &model.Wallet{}
	result := db.First(wallet, "id = ? || uuid = ?", id, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_FOUND)
	}
	return wallet, nil
}

func CreateWallet(wallet *model.Wallet, db *gorm.DB) (*model.Wallet, error){
	helper.Logger().Info("Creating wallet for " + wallet.EntityType + " with uuid " + wallet.EntityUuid)
	result := db.Create(wallet)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_INSERTED)
	}
	return wallet, nil
}

func UpdateWalletIsActive(id string, isActive bool, db *gorm.DB) (*model.Wallet, error) {
	helper.Logger().Info("Fetching wallet for id" + id)
	wallet := &model.Wallet{}
	result := db.First(wallet, "id = ? || uuid = ?", id, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_FOUND)
	}
	wallet.IsActive = isActive
	result = db.Model(wallet).Update("is_active", isActive)
	if result.Error != nil {
		return nil, result.Error
	}
	return wallet, nil
}

func GetEntityWallet(entityUuid string, entityType string, db *gorm.DB) (*model.Wallet, error) {
	wallet := &model.Wallet{}
	result := db.First(wallet, "entity_uuid = ? and entity_type = ?", entityUuid, entityType)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(config.DATA_NOT_FOUND)
	}

	return wallet, nil
}

func AreWalletsPresent(ids []string, db *gorm.DB) ([]*dto.GetIdDto, error) {
	fmt.Printf("Searching wallet with ids %v", ids)
	var getIds []*dto.GetIdDto
	db.Find(&getIds, "uuid in ?", ids)
	if len(getIds) != len(ids) {
		return nil, errors.New(config.WALLETS_IN_TRANSACTION_NOT_AVAILABLE)
	}
	return getIds,nil
}

func AddToWallet(id int64, amount float64, db *gorm.DB) error {
	result := db.Find(&model.Wallet{}, id).Update("current_balance", gorm.Expr("current_balance + ?", 1, amount))
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SubFromWallet(id int64, amount float64, db *gorm.DB) error {
	result := db.Model(&model.Wallet{}).Where("id = ? and current_balance > ", id, amount).Update("current_balance", gorm.Expr("current_balance - ?", 1, amount))
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New(config.UNABLE_TO_DEBIT_FROM_WALLET)
	}
	return nil
}
