package router

import (
	"github.com/aqilknz/shortlink-backend/internal/controller"
	"github.com/aqilknz/shortlink-backend/internal/middleware"
	"github.com/aqilknz/shortlink-backend/internal/repository"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, authController *controller.AuthController, authRepo repository.AuthRepository) {
	authGroup := rg.Group("/auth")
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
	authGroup.DELETE("/logout", middleware.RequireAuth(authRepo), authController.Logout)
}
