package services

import (
	"github.com/ofojichigozie/iphms-go-backend/dtos"
	"github.com/ofojichigozie/iphms-go-backend/models"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
)

type VitalsService interface {
	CreateVitals(input dtos.CreateVitalsInput) (models.Vitals, error)
	GetAllVitals(filters map[string]interface{}) ([]models.Vitals, error)
	GetVitalsById(vitalsId uint) (models.Vitals, error)
	DeleteVitalsById(vitalsId uint) error
}

type vitalsService struct {
	vitalsRepository repositories.VitalsRepository
	userRepository   repositories.UserRepository
}

func NewVitalsService(vitalsRepository repositories.VitalsRepository,
	userRepository repositories.UserRepository) VitalsService {
	return &vitalsService{vitalsRepository: vitalsRepository, userRepository: userRepository}
}

func (s *vitalsService) CreateVitals(input dtos.CreateVitalsInput) (models.Vitals, error) {
	_, err := s.userRepository.FindByID(input.UserId)
	if err != nil {
		return models.Vitals{}, err
	}

	vitals := models.Vitals{
		Temperature:    input.Temperature,
		Humidity:       input.Humidity,
		PulseRate:      input.PulseRate,
		LightIntensity: input.LightIntensity,
		UserID:         input.UserId,
	}

	if err := s.vitalsRepository.Create(&vitals); err != nil {
		return models.Vitals{}, err
	}

	return vitals, nil
}

func (s *vitalsService) GetAllVitals(filters map[string]interface{}) ([]models.Vitals, error) {
	return s.vitalsRepository.FindAll(filters)
}

func (s *vitalsService) GetVitalsById(vitalsId uint) (models.Vitals, error) {
	return s.vitalsRepository.FindByID(vitalsId)
}

func (s *vitalsService) DeleteVitalsById(vitalsId uint) error {
	return s.vitalsRepository.Delete(vitalsId)
}
