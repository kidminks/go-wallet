package setup

import (
	"github.com/kidminks/go_wallet/setup"
	"github.com/kidminks/go_wallet/test/database"
	"testing"
)

func TestSetup(t *testing.T) {
	db := database.GetDbConnection()
	setup.Setup(db)
}


