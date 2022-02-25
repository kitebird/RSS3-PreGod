package db

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() error {
	err := mgm.SetDefaultConfig(nil, config.Config.Mongo.DB, options.Client().ApplyURI(config.Config.Mongo.URI))
	// TODO: Create Indexes
	// Reference: https://github.com/Kamva/mgm/issues/35
	return err
}
