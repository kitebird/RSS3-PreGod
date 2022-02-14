package model

import "github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"

// `account_platform` model.
type AccountPlatform struct {
	Base BaseModel `gorm:"embedded"`

	AccountID         string                   `gorm:"primaryKey;type:text;column:account_id"`
	PlatformName      constants.PlatformNameID `gorm:"type:int"`
	PlatformAccountID string                   `gorm:"type:text"` // account ID on the platform
}
