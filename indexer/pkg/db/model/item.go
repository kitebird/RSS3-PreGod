package model

import (
	"time"

	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"github.com/kamva/mgm/v3"
)

type ItemId struct {
	Tags  constants.ItemTags `json:"tags" bson:"tags"`
	Proof string             `json:"proof" bson:"proof"`
}
type Item struct {
	mgm.DefaultModel `bson:",inline"`

	ItemId ItemId `json:"item_id" bson:"item_id"` // Index

	ObjectUid string `json:"object_uid" bson:"object_uid"`

	From              string    `json:"from" bson:"from"`
	To                string    `json:"to" bson:"to"`
	PlatformCreatedAt time.Time `json:"date_created" bson:"date_created"`
}

func NewItem(objectUid string, tags constants.ItemTags, from string, to string, proof string, platformCreatedAt time.Time) *Item {
	return &Item{
		ItemId: ItemId{
			Tags:  tags,
			Proof: proof,
		},

		ObjectUid: objectUid,

		From:              from,
		To:                to,
		PlatformCreatedAt: platformCreatedAt,
	}
}
