package model

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Item struct {
	mgm.DefaultModel `bson:",inline"`

	ObjectUid  string               `json:"object_uid" bson:"object_uid"`
	ObjectType constants.ItemTypeID `json:"object_type:int" bson:"object_type"`
	Proof      string               `json:"proof" bson:"proof"` // Indexes: (Proof, ObjectType)

	From              string    `json:"from" bson:"from"`
	To                string    `json:"to" bson:"to"`
	PlatformCreatedAt time.Time `json:"date_created" bson:"date_created"`
}

func NewItem(objectUid string, objectType constants.ItemTypeID, from string, to string, proof string, platformCreatedAt time.Time) *Item {
	return &Item{
		ObjectUid:  objectUid,
		ObjectType: objectType,
		Proof:      proof,

		From:              from,
		To:                to,
		PlatformCreatedAt: platformCreatedAt,
	}
}

func InsertItemDoc(item *Item) *mongo.SingleResult {
	return mgm.Coll(&Item{}).FindOneAndReplace(mgm.Ctx(), bson.M{"object_type": item.ObjectType, "proof": item.Proof}, item, options.FindOneAndReplace().SetUpsert(true))
}

// TODO: getter
