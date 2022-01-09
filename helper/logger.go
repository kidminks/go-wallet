package helper

import (
"fmt"
"go.uber.org/zap"
"sync"
)

var loggerLock = &sync.Mutex{}

var logger *zap.Logger

func getLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

func Logger() *zap.Logger {
	if logger == nil {
		loggerLock.Lock()
		defer loggerLock.Unlock()
		if logger == nil {
			fmt.Println("Creating single instance for logger.")
			logger = getLogger()
		}
	}
	return logger
}
