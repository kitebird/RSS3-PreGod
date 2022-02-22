package data_migration

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"log"
)

func prepareDB() error {
	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup err: %v", err)
	}

	if err := db.Setup(); err != nil {
		log.Fatalf("db.Setup err: %v", err)
	}

	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("db.AutoMigrate err: %v", err)
	}

	return nil
}
