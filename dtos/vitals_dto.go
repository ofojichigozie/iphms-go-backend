package dtos

type CreateVitalsInput struct {
	Temperature    float32 `json:"temperature" binding:"required"`
	Humidity       float32 `json:"humidity" binding:"required"`
	PulseRate      float32 `json:"pulseRate" binding:"required"`
	LightIntensity float32 `json:"lightIntensity" binding:"required"`
	UserId         uint    `json:"userId" binding:"required"`
}
