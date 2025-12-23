package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/controllers"
	"github.com/yash-sojitra-20/address-book-backend/internal/db"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), middleware.GlobalRecovery())

	userRepo := repositories.NewUserRepository(db.DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	return r
}
