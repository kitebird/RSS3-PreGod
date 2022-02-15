package model

// `asset` model.
type Asset struct {
	Base BaseModel `gorm:"embedded"`

	AssetID string `gorm:"primaryKey;type:text;column:asset_id"`
	ItemID  string `gorm:"type:text;column:item_id"`
}
