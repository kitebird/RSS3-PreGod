package model

import "github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"

// `account_platform` model.
type AccountPlatform struct {
	Base BaseModel `gorm:"embedded"`

	AccountID         string               `gorm:"primaryKey;type:text;column:account_id"`
	PlatformNameID    constants.PlatformID `gorm:"type:int;column:platform_name_id"`
	PlatformAccountID string               `gorm:"type:text;column:platform_account_id"` // account ID on the platform
}
