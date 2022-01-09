package config

import "taskpot.com/go_payments/model/dto"

const CreateTableSqlPath = "./setup/config/database/sql_table_script.sql"
const MySql = "mysql"
var ParentDbConf dto.ParentDatabaseConfiguration

func SetParentDatabaseConfig(configuration dto.ParentDatabaseConfiguration) {
	ParentDbConf = configuration
}

