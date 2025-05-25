package models

import (
	"time"
)

type Merchant struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	Ratings   int        `gorm:"not null;default:0" json:"ratings"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}

// Optional: custom table name explicit kalau mau
func (Merchant) TableName() string {
	return "merchants"
}
