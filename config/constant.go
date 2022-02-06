package config

import (
	"github.com/kidminks/go_wallet/test/database/dto"
)

const MySql = "mysql"
var ParentDbConf dto.ParentDatabaseConfiguration

func SetParentDatabaseConfig(configuration dto.ParentDatabaseConfiguration) {
	ParentDbConf = configuration
}

