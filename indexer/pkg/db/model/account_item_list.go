package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountItemList struct {
	mgm.DefaultModel `bson:",inline"`

	AccountInstance string `json:"account_instance" bson:"account_instance"`

	Assets []ItemId `json:"assets" bson:"assets"`

	Notes []ItemId `json:"notes" bson:"notes"`
}

func SetAssets(instance string, assets []ItemId) {
	mgm.Coll(&AccountItemList{}).FindOneAndUpdate(
		mgm.Ctx(), bson.M{"account_instance": instance},
		bson.M{"$set": bson.M{"assets": assets}},
		options.FindOneAndUpdate().SetUpsert(true),
	)
}

func AppendNotes(instance string, notes []ItemId) {
	// If the value is a document, MongoDB determines that the document is a duplicate if an existing
	// document in the array matches the to-be-added document exactly; i.e. the existing document has
	// the exact same fields and values and the fields are in the same order. As such, field order
	// matters and you cannot specify that MongoDB compare only a subset of the fields in the document
	// to determine whether the document is a duplicate of an existing array element.
	// If we use ODM, the order is the same. So we do not need to worry
	mgm.Coll(&AccountItemList{}).FindOneAndUpdate(
		mgm.Ctx(),
		bson.M{"account_instance": instance}, bson.M{"$addToSet": bson.M{"notes": bson.M{"$each": notes}}},
		options.FindOneAndUpdate().SetUpsert(true),
	)
}

// TODO: getter
