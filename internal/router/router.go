package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/controllers"
	"github.com/yash-sojitra-20/address-book-backend/internal/db"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
	"github.com/yash-sojitra-20/address-book-backend/internal/repositories"
	"github.com/yash-sojitra-20/address-book-backend/internal/services"
)

var cfg *config.Config

func SetConfig(c *config.Config) {
	cfg = c
	// fmt.Println(jwtSecret)
}

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

	addressRepo := repositories.NewAddressRepository(db.DB)
	addressService := services.NewAddressService(addressRepo)
	addressController := controllers.NewAddressController(addressService, cfg)

	addresses := r.Group("/addresses")
	addresses.Use(middleware.AuthMiddleware())
	addresses.Use(middleware.EnsureUserExistsMiddleware(userRepo))
	{
		addresses.POST("", addressController.Create)
		addresses.GET("", addressController.GetAll)
		addresses.GET("/:id", addressController.GetByID)
		addresses.GET("/export/custom", addressController.ExportCustom)
		addresses.GET("/export/sync", addressController.Export)
		addresses.GET("/export/async", addressController.ExportAsync)
		addresses.PUT("/:id", addressController.Update)
		addresses.DELETE("/:id", addressController.Delete)
	}

	return r
}
