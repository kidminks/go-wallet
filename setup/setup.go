package setup

import (
	"gorm.io/gorm"
	"taskpot.com/go_wallet/helper"
)

func Setup(db *gorm.DB) {
	helper.Logger().Info("Golang wallet library setup")
	databaseSetup(db)
}
