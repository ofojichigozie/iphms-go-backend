package services

import (
	"errors"
	"strings"
	"time"

	"github.com/ofojichigozie/iphms-go-backend/dtos"
	"github.com/ofojichigozie/iphms-go-backend/models"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
	"github.com/ofojichigozie/iphms-go-backend/utils"
)

type UserService interface {
	CreateUser(input dtos.CreateUserInput) (models.User, error)
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByDeviceID(deviceID string) (models.User, error)
	UpdateUser(id uint, input dtos.UpdateUserInput) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (us *userService) CreateUser(input dtos.CreateUserInput) (models.User, error) {
	if _, err := time.Parse("2006-01-02", input.DateOfBirth); err != nil {
		return models.User{}, errors.New("invalid date format. Use YYYY-MM-DD format")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:        input.Name,
		Email:       strings.ToLower(input.Email),
		Password:    hashedPassword,
		DateOfBirth: input.DateOfBirth,
		DeviceId:    input.DeviceId,
		Role:        "user",
	}

	if err := us.userRepository.Create(&user); err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (us *userService) GetUsers() ([]models.User, error) {
	return us.userRepository.FindAll()
}

func (us *userService) GetUserByID(id uint) (models.User, error) {
	return us.userRepository.FindByID(id)
}

func (us *userService) GetUserByEmail(email string) (models.User, error) {
	return us.userRepository.FindByEmail(email)
}

func (us *userService) GetUserByDeviceID(deviceID string) (models.User, error) {
	return us.userRepository.FindByDeviceID(deviceID)
}

func (us *userService) UpdateUser(id uint, input dtos.UpdateUserInput) (models.User, error) {
	user, err := us.userRepository.FindByID(id)
	if err != nil {
		return models.User{}, err
	}

	updates := make(map[string]interface{})

	if input.Name != nil {
		updates["name"] = *input.Name
		user.Name = *input.Name
	}

	if input.Email != nil {
		lowerEmail := strings.ToLower(*input.Email)
		updates["email"] = lowerEmail
		user.Email = lowerEmail
	}

	if input.Password != nil {
		hashedPassword, err := utils.HashPassword(*input.Password)
		if err != nil {
			return user, err
		}
		updates["password"] = hashedPassword
		user.Password = hashedPassword
	}

	if input.DateOfBirth != nil {
		updates["date_of_birth"] = *input.DateOfBirth
		user.DateOfBirth = *input.DateOfBirth
	}

	if input.DeviceId != nil {
		updates["device_id"] = *input.DeviceId
		user.DeviceId = *input.DeviceId
	}

	if len(updates) == 0 {
		return user, errors.New("no valid fields provided for update")
	}

	if err := us.userRepository.Update(&user, updates); err != nil {
		return user, err
	}

	return user, nil
}

func (us *userService) DeleteUser(id uint) error {
	return us.userRepository.Delete(id)
}
