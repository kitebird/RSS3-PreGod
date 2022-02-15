package model

// `link` model.
type Link struct {
	Base BaseModel `gorm:"embedded"`

	LinkID string `gorm:"primaryKey;type:text;column:asset_id"`
	ItemID string `gorm:"type:text;column:item_id"`

	RSS3ID       string `gorm:"type:text;column:rss3_id"`
	TargetRSS3ID string `gorm:"type:text;column:target_rss3_id"`

	PageIndex int `gorm:"type:int"`
}
