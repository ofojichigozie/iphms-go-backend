package initializers

import (
	"fmt"
	"log"
	"time"

	"github.com/ofojichigozie/iphms-go-backend/models"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
	"github.com/ofojichigozie/iphms-go-backend/utils"
)

func InitAdminUser() {
	userRepository := repositories.NewUserRepository(DB)

	adminEmail := "system.admin@iphms.dev"
	_, err := userRepository.FindByEmail(adminEmail)
	if err != nil {
		hashedPassword, err := utils.HashPassword("admin12345")
		if err != nil {
			log.Printf("Error hashing admin password: %v", err)
			return
		}

		adminUser := models.User{
			Name:        "System Admin",
			Email:       adminEmail,
			Password:    hashedPassword,
			DateOfBirth: time.Now().Format("2006-01-02"),
			DeviceId:    "system-admin-device",
			Role:        "admin",
		}

		if err := userRepository.Create(&adminUser); err != nil {
			log.Printf("Error creating admin user: %v", err)
			return
		}

		fmt.Println("Default admin user created successfully")
	} else {
		fmt.Println("Admin user already exists, skipping creation")
	}
}
