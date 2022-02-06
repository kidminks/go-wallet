package setup

import (
	"gorm.io/gorm"
	"github.com/kidminks/go_wallet/helper"
)

func Setup(db *gorm.DB) {
	helper.Logger().Info("Golang wallet library setup")
	databaseSetup(db)
}
