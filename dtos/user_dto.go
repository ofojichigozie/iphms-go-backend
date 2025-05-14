package dtos

type CreateUserInput struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required"`
	DateOfBirth string `json:"dateOfBirth" binding:"required"`
	DeviceId    string `json:"deviceId" binding:"required"`
}

type UpdateUserInput struct {
	Name        *string `json:"name"`
	Email       *string `json:"email" binding:"omitempty,email"`
	Password    *string `json:"password"`
	DateOfBirth *string `json:"dateOfBirth"`
	DeviceId    *string `json:"deviceId"`
}
