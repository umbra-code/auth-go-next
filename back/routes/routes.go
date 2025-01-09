package routes

import (
	"api-auth/internal/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Ruta pública para autenticación
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", auth.LoginHandler)
		authGroup.POST("/register", auth.RegisterHandler)
		authGroup.POST("/refresh", auth.RefreshHandler)
	}

	// Ruta protegida por middleware
	protectedGroup := router.Group("/user")
	protectedGroup.Use(auth.AuthMiddleware())
	{
		protectedGroup.GET("/profile", auth.UserProfileHandler)
	}
}
