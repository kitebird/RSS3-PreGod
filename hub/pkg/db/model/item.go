package model

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

// `item` model.
type Item struct {
	Base BaseModel `gorm:"embedded"`

	ItemID string `gorm:"primaryKey;type:text;column:item_id"`

	ObjectID string `gorm:"type:text"` // TODO: association

	Proof     string               `gorm:"type:text"`
	Type      constants.ItemTypeID `gorm:"type:int"`
	From      string               `gorm:"type:text"`
	To        string               `gorm:"type:text"`
	Auto      bool                 `gorm:"type:bool"`
	PageIndex int                  `gorm:"type:int"`
}
