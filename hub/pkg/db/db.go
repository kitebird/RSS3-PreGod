package db

import (
	"time"

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
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})

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
