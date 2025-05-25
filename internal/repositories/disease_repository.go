package repositories

import (
	"fmt"

	"github.com/nix-united/golang-echo-boilerplate/internal/models"
	"gorm.io/gorm"
)

type DiseaseRepository interface {
	GetDiseasesByUser(userId int) ([]models.MedicalDisease, error)
	Create(diseases *[]models.Disease) error
}

type diseaseRepository struct {
	db *gorm.DB
}

func NewDiseaseRepository(db *gorm.DB) DiseaseRepository {
	return &diseaseRepository{db: db}
}

func (r diseaseRepository) GetDiseasesByUser(userId int) ([]models.MedicalDisease, error) {
	var diseases []models.MedicalDisease
	if err := r.db.Preload("User").Preload("Disease").Where("user_id = ?", userId).Find(&diseases).Error; err != nil {
		return nil, fmt.Errorf("execute select posts query: %w", err)
	}

	return diseases, nil
}

func (r diseaseRepository) Create(diseases *[]models.Disease) error {
	if err := r.db.Create(diseases).Error; err != nil {
		return fmt.Errorf("execute insert diseases query: %w", err)
	}

	return nil
}
