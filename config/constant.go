package config

import (
	"taskpot.com/go_wallet/test/database/dto"
)

const MySql = "mysql"
var ParentDbConf dto.ParentDatabaseConfiguration

func SetParentDatabaseConfig(configuration dto.ParentDatabaseConfiguration) {
	ParentDbConf = configuration
}

