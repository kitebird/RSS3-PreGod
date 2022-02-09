package logger

import (
	"log"

	"go.uber.org/zap"
)

func InitSugaredLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync() // ignore error (if any)
	sugar := logger.Sugar()
	return sugar
}
