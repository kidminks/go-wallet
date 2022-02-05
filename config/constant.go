package config

import "taskpot.com/go_wallet/model/dto"

const MySql = "mysql"
var ParentDbConf dto.ParentDatabaseConfiguration

func SetParentDatabaseConfig(configuration dto.ParentDatabaseConfiguration) {
	ParentDbConf = configuration
}

