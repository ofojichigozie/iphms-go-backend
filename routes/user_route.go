package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/controllers"
	"github.com/ofojichigozie/iphms-go-backend/middleware"
	"github.com/ofojichigozie/iphms-go-backend/repositories"
	"github.com/ofojichigozie/iphms-go-backend/services"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.Engine, DB *gorm.DB) {
	userRepository := repositories.NewUserRepository(DB)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	userGroup := r.Group("/users")
	{
		userGroup.POST("", userController.CreateUser)
		userGroup.Use(middleware.AuthMiddleware())
		{
			userGroup.GET("", userController.GetUsers)
			userGroup.GET("/:id", userController.GetUser)
			userGroup.PATCH("/:id", userController.UpdateUser)
			userGroup.DELETE("/:id", userController.DeleteUser)

		}
	}
}
