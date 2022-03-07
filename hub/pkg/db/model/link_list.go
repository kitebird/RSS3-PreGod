package model

import (
	datatype "github.com/NaturalSelectionLabs/RSS3-PreGod/hub/pkg/db/model/datatype"
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

// `link_list` model.
type LinkList struct {
	Base BaseModel `gorm:"embedded"`

	LinkListID string `gorm:"primaryKey;type:text;column:link_list_id"`

	RSS3ID string `gorm:"type:text;column:rss3_id"` // owner id

	LinkType constants.LinkTypeID `gorm:"type:int"`

	Metadata datatype.Attachments `gorm:"type:jsonb"`
}
