package model

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ItemId struct {
	ItemType constants.ItemTypeID `json:"item_type:int" bson:"item_type"`
	Proof    string               `json:"proof" bson:"proof"`
}
type Item struct {
	mgm.DefaultModel `bson:",inline"`

	ItemId ItemId `json:"item_id" bson:"item_id"` // Index

	ObjectUid string `json:"object_uid" bson:"object_uid"`

	From              string    `json:"from" bson:"from"`
	To                string    `json:"to" bson:"to"`
	PlatformCreatedAt time.Time `json:"date_created" bson:"date_created"`
}

func NewItem(objectUid string, objectType constants.ItemTypeID, from string, to string, proof string, platformCreatedAt time.Time) *Item {
	return &Item{
		ItemId: ItemId{
			ItemType: objectType,
			Proof:    proof,
		},

		ObjectUid: objectUid,

		From:              from,
		To:                to,
		PlatformCreatedAt: platformCreatedAt,
	}
}

func InsertItemDoc(item *Item) *mongo.SingleResult {
	return mgm.Coll(&Item{}).FindOneAndReplace(
		mgm.Ctx(),
		bson.M{"item_id.item_type": item.ItemId.ItemType, "item_id.proof": item.ItemId.Proof},
		item,
		options.FindOneAndReplace().SetUpsert(true),
	)
}

// TODO: getter
