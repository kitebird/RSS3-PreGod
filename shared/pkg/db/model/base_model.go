package model

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	Version   string `gorm:"column:version;type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
