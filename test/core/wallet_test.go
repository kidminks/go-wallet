package core

import (
	"errors"
	"github.com/kidminks/go_wallet/core"
	"github.com/kidminks/go_wallet/test/database"
	"github.com/kidminks/go_wallet/test/database/util"
	"testing"
)

func TestCreateWallet(t *testing.T) {
	wallet := util.GetBasicWallet("user_id_1")
	db := database.GetDbConnection()
	wallet, err := core.CreateWallet(wallet, db)
	if err != nil {
		panic(err)
	}
}

func TestGetWalletById(t *testing.T) {
	db := database.GetDbConnection()
	_, err := core.GetWalletById("1", db)
	if err != nil {
		panic(err)
	}
}

func TestUnableToGetError(t *testing.T) {
	db := database.GetDbConnection()
	_, err := core.GetWalletById("10", db)
	if err == nil {
		panic(errors.New("panic should come"))
	}
}

func TestUpdateWalletIsActive(t *testing.T) {
	db := database.GetDbConnection()
	_, err := core.UpdateWalletIsActive("1", true, db)
	if err != nil {
		panic(err)
	}
}

func TestGetEntityWallet(t *testing.T) {
	db := database.GetDbConnection()
	_, err := core.GetEntityWallet("user_id_1", "USER", db)
	if err != nil {
		panic(err)
	}
}