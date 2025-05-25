package models

import (
	"time"
)

type MedicalAllergy struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64      `gorm:"not null" json:"user_id"`
	AllergyID int64      `gorm:"not null" json:"allergy_id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional relations
	User    *User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Allergy *Allergy `gorm:"foreignKey:AllergyID" json:"allergy,omitempty"`
}
