package db

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() error {
	// Establish a connection to the database
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: config.Postgres.DSN,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	// Ping
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	defer sqlDB.Close()

	if err = sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}
