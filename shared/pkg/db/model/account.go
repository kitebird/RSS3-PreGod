package model

import "github.com/lib/pq"

// `account` model.
type Account struct {
	Base BaseModel `gorm:"embedded"`

	AccountID string         `gorm:"primaryKey;type:text;column:account_id"`
	Name      string         `gorm:"type:text"`
	Bio       string         `gorm:"type:text"`
	Avatars   pq.StringArray `gorm:"type:text[]"`

	// The following fields are stored in `attachments` field:
	// Banners   pq.StringArray `gorm:"type:text[]"`
	// Websites  pq.StringArray `gorm:"type:text[]"`

	InstanceBase    InstanceBase
	AccountPlatform []AccountPlatform
}
