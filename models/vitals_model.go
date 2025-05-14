package models

import (
	"gorm.io/gorm"
)

type Vitals struct {
	gorm.Model
	Temperature    float32 `json:"temperature" gorm:"not null"`
	Humidity       float32 `json:"humidity" gorm:"not null"`
	PulseRate      float32 `json:"pulseRate" gorm:"not null"`
	LightIntensity float32 `json:"lightIntensity" gorm:"not null"`
	UserID         uint    `json:"userId" gorm:"not null"`
	User           User    `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
