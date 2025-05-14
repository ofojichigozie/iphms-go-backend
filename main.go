package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ofojichigozie/iphms-go-backend/initializers"
	"github.com/ofojichigozie/iphms-go-backend/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.InitAdminUser()
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://iphms-app.onrender.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With", "X-Device-ID", "X-Device-Secret"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.AuthRoute(r, initializers.DB)
	routes.UserRoutes(r, initializers.DB)
	routes.VitalsRoutes(r, initializers.DB)
	r.Run()
}
