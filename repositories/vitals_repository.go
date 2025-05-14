package repositories

import (
	"errors"

	"github.com/ofojichigozie/iphms-go-backend/models"
	"gorm.io/gorm"
)

type VitalsRepository interface {
	Create(vitals *models.Vitals) error
	FindAll(filters map[string]interface{}) ([]models.Vitals, error)
	FindByID(id uint) (models.Vitals, error)
	Delete(id uint) error
}

type vitalsRepository struct {
	db *gorm.DB
}

func NewVitalsRepository(db *gorm.DB) VitalsRepository {
	return &vitalsRepository{db: db}
}

func (r *vitalsRepository) Create(vitals *models.Vitals) error {
	return r.db.Create(vitals).Error
}

func (r *vitalsRepository) FindAll(filters map[string]interface{}) ([]models.Vitals, error) {
	var vitals []models.Vitals
	query := r.db.Model(&models.Vitals{})

	for key, value := range filters {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&vitals).Error; err != nil {
		return nil, err
	}

	return vitals, nil
}

func (r *vitalsRepository) FindByID(id uint) (models.Vitals, error) {
	var vitals models.Vitals
	if err := r.db.First(&vitals, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return vitals, errors.New("vitals not found")
		}
		return vitals, err
	}
	return vitals, nil
}

func (r *vitalsRepository) Delete(id uint) error {
	return r.db.Unscoped().Delete(&models.Vitals{}, id).Error
}
