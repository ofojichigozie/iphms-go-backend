package repositories

import (
	"errors"
	"strings"

	"github.com/ofojichigozie/iphms-go-backend/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (models.User, error)
	FindByEmail(email string) (models.User, error)
	FindByDeviceID(deviceID string) (models.User, error)
	Update(user *models.User, updates map[string]interface{}) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", strings.ToLower(email)).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, err
	}
	return user, nil
}

func (r *userRepository) FindByDeviceID(deviceID string) (models.User, error) {
	var user models.User
	if err := r.db.Where("device_id = ?", deviceID).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User, updates map[string]interface{}) error {
	return r.db.Model(user).Updates(updates).Error
}

func (r *userRepository) Delete(id uint) error {
	// Start a transaction
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Delete associated vitals records first
	if err := tx.Where("user_id = ?", id).Unscoped().Delete(&models.Vitals{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Then delete the user
	if err := tx.Unscoped().Delete(&models.User{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
