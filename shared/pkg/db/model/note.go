package model

// `note` model.
type Note struct {
	Base BaseModel `gorm:"embedded"`

	NoteID string `gorm:"primaryKey;type:text;column:note_id"`
	ItemID string `gorm:"type:text"` // TODO: association
}
