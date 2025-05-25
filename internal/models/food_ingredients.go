package models

import (
	"time"
)

type FoodIngredient struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	FoodID    int64      `gorm:"not null" json:"food_id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`

	// Optional relation biar gampang preload
	Food *MerchantFood `gorm:"foreignKey:FoodID" json:"food,omitempty"`
}

// Optional custom table name
func (FoodIngredient) TableName() string {
	return "food_ingredients"
}
