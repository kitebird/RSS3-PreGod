package model

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
	"gorm.io/datatypes"
)

// `instance_base` model.
type InstanceBase struct {
	Base BaseModel `gorm:"embedded"`

	UUID              string                   `gorm:"primaryKey;type:text;column:uuid"`
	ControllerAddress string                   `gorm:"type:text"`
	InstanceType      constants.InstanceTypeID `gorm:"type:int"`
	Attachments       datatypes.JSON           `gorm:"type:jsonb"`
}
