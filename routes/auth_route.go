package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/controllers"
	"github.com/ofojichigozie/iphms-go-backend/middleware"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
	"github.com/ofojichigozie/iphms-go-backend/services"
	"gorm.io/gorm"
)

func AuthRoute(r *gin.Engine, DB *gorm.DB) {
	userRepository := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(userService)

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
		authGroup.Use(middleware.AuthMiddleware())
		{
			authGroup.POST("/refresh", authController.RefreshToken)
		}
	}
}
