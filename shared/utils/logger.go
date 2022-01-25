package utils

import "go.uber.org/zap"

func InitSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // ignore error (if any)
	sugar := logger.Sugar()
	return sugar
}
