package utils

import (
	"encoding/json"
	"os"
	"taskpot.com/go_wallet/config"
	"taskpot.com/go_wallet/test/database/dto"
)

func ReadParentDatabaseConfig(filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := dto.ParentDatabaseConfiguration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
	config.SetParentDatabaseConfig(configuration)
}

