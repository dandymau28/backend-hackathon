package repositories

import (
	"fmt"

	"github.com/nix-united/golang-echo-boilerplate/internal/models"
	"gorm.io/gorm"
)

type AllergyRepository interface {
	GetAllergiesByUser(userId int) ([]models.MedicalAllergy, error)
	Create(allergies *[]models.Allergy) error
}

type allergyRepository struct {
	db *gorm.DB
}

func NewAllergyRepository(db *gorm.DB) AllergyRepository {
	return &allergyRepository{db: db}
}

func (r *allergyRepository) GetAllergiesByUser(userId int) ([]models.MedicalAllergy, error) {
	var allergies []models.MedicalAllergy
	if err := r.db.Preload("User").Preload("Allergy").Where("user_id = ?", userId).Find(&allergies).Error; err != nil {
		return nil, fmt.Errorf("execute select posts query: %w", err)
	}

	return allergies, nil
}

func (r *allergyRepository) Create(allergies *[]models.Allergy) error {
	if err := r.db.Create(allergies).Error; err != nil {
		return fmt.Errorf("execute insert allergies query: %w", err)
	}

	return nil
}
