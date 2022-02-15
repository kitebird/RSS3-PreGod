package model

import (
	"gorm.io/datatypes"
)

// `third_party_storage` model.
type Third_Party_Storage struct {
	Base BaseModel `gorm:"embedded"`

	StorageID string `gorm:"primaryKey;type:text;column:storage_id"`
	RSS3ID    string `gorm:"type:text"`

	AppName     string         `gorm:"type:text"`
	Key         string         `gorm:"type:text"`
	Value       datatypes.JSON `gorm:"type:jsonb"`
	Description string         `gorm:"type:text"`
}
