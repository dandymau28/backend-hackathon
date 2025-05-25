package models

import (
	"time"
)

type Allergy struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"` // soft delete friendly
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
