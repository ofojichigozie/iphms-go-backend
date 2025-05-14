package main

import (
	"fmt"

	"github.com/ofojichigozie/iphms-go-backend/initializers"
	"github.com/ofojichigozie/iphms-go-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	// Drop existing foreign key constraints first to avoid conflicts
	initializers.DB.Exec("ALTER TABLE vitals DROP CONSTRAINT IF EXISTS fk_vitals_user")

	// Migrate the models
	initializers.DB.AutoMigrate(&models.User{}, &models.Vitals{})

	// Ensure the foreign key with cascade delete is properly set up
	initializers.DB.Exec("ALTER TABLE vitals ADD CONSTRAINT fk_vitals_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE")

	fmt.Println("Database migration successful with cascade delete constraints")
}
