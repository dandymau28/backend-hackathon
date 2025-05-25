package models

import (
	"time"
)

type MedicalDisease struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    int64      `gorm:"not null" json:"user_id"`
	DiseaseID int64      `gorm:"not null" json:"disease_id"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional relations
	User    *User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Disease *Disease `gorm:"foreignKey:DiseaseID" json:"disease,omitempty"`
}

// Optional: kalau mau custom table name explicit
func (MedicalDisease) TableName() string {
	return "medical_diseases"
}
