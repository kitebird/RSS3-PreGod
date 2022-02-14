package db

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func Setup() error {
	// Establish a connection to the database
	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: config.Postgres.DSN,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
	})

	if err != nil {
		return err
	}

	return nil
}
