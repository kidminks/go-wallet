package core

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/database/dto"
	"taskpot.com/go_wallet/helper"
	"taskpot.com/go_wallet/model"
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

func AreWalletsPresent(ids []string, db *gorm.DB) ([]*dto.GetIdDto, error) {
	fmt.Printf("Searching wallet with ids %v", ids)
	var getIds []*dto.GetIdDto
	db.Find(&getIds, "uuid in ?", ids)
	if len(getIds) != len(ids) {
		return nil, errors.New(config.WALLETS_IN_TRANSACTION_NOT_AVAILABLE)
	}
	return getIds,nil
}
