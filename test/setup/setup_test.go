package setup

import (
	"taskpot.com/go_wallet/setup"
	"taskpot.com/go_wallet/test/database"
	"testing"
)

func TestSetup(t *testing.T) {
	db := database.GetDbConnection()
	setup.Setup(db)
}


