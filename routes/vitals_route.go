package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/controllers"
	"github.com/ofojichigozie/iphms-go-backend/middleware"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
	"github.com/ofojichigozie/iphms-go-backend/services"
	"gorm.io/gorm"
)

func VitalsRoutes(router *gin.Engine, DB *gorm.DB) {
	userRepository := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepository)
	vitalRepository := repositories.NewVitalsRepository(DB)
	vitalsService := services.NewVitalsService(vitalRepository, userRepository)
	vitalsController := controllers.NewVitalsController(vitalsService)

	vitalsGroup := router.Group("/vitals")
	{
		vitalsGroup.POST("",
			middleware.IoTDeviceMiddleware(userService),
			vitalsController.CreateVitals)
		vitalsGroup.Use(middleware.AuthMiddleware())
		{
			vitalsGroup.GET("", vitalsController.GetAllVitals)
			vitalsGroup.GET("/:id", vitalsController.GetVitals)
			vitalsGroup.DELETE("/:id", vitalsController.DeleteVital)
		}
	}
}
