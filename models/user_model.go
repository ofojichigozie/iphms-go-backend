package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Email       string `json:"email" gorm:"unique;not null"`
	Password    string `json:"-" gorm:"not null"`
	DateOfBirth string `json:"dateOfBirth" gorm:"type:date;not null"`
	DeviceId    string `json:"deviceId" gorm:"not null;unique"`
	Role        string `json:"role" gorm:"not null;default:user"`
}
