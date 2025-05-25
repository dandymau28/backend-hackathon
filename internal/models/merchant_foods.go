package models

import (
	"time"
)

type MerchantFood struct {
	ID         int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	MerchantID int64      `gorm:"not null" json:"merchant_id"`
	Name       string     `gorm:"type:varchar(255);not null" json:"name"`
	Price      int        `gorm:"not null;default:0" json:"price"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt  time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional: relasi ke Merchant biar gampang preload
	Merchant *Merchant `gorm:"foreignKey:MerchantID" json:"merchant,omitempty"`
}

// Optional custom table name
func (MerchantFood) TableName() string {
	return "merchant_foods"
}
