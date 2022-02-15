package model

// `link` model.
type Link struct {
	Base BaseModel `gorm:"embedded"`

	LinkID string `gorm:"primaryKey;type:text;column:asset_id"`
	ItemID string `gorm:"type:text"`

	RSS3ID       string `gorm:"type:text"`
	TargetRSS3ID string `gorm:"type:text"`

	PageIndex int `gorm:"type:int"`
}
