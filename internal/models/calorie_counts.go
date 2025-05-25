package models

import (
	"time"
)

type CalorieCount struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64      `gorm:"not null" json:"user_id"`
	Datetime  time.Time  `gorm:"type:timestamptz;not null" json:"datetime"`
	Calories  int        `gorm:"not null;default:0" json:"calories"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional relation biar preload data user
	User *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Optional custom table name
func (CalorieCount) TableName() string {
	return "calorie_counts"
}
