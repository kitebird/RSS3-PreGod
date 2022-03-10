package main

import (
	"log"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/router"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/cache"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/logger"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/web"
)

func init() {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := logger.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := cache.Setup(); err != nil {
		logger.Fatalf("cache.Setup err: %v", err)
	}

	if err := db.Setup(); err != nil {
		logger.Fatalf("db.Setup err: %v", err)
	}

	if err := db.AutoMigrate(); err != nil {
		logger.Fatalf("db.AutoMigrate err: %v", err)
	}
}

func main() {
	addr := web.Setup(&web.Config{
		RunMode:      config.Config.HubServer.RunMode,
		HttpPort:     config.Config.HubServer.HttpPort,
		ReadTimeout:  config.Config.HubServer.ReadTimeout,
		WriteTimeout: config.Config.HubServer.WriteTimeout,
		Handler:      router.InitRouter(),
	})

	logger.Infof("Start http server listening on http://%s", addr)
	defer logger.Logger.Sync()
}
