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
func SetConfig(c *config.Config){
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

	contactRepo := repositories.NewContactRepository(db.DB)
	contactService := services.NewContactService(contactRepo)
	contactController := controllers.NewContactController(contactService, cfg)

	contacts := r.Group("/contacts")
	contacts.Use(middleware.AuthMiddleware())
	{
		contacts.POST("", contactController.Create)
		contacts.GET("", contactController.GetAll)
		contacts.GET("/export", contactController.Export)
		contacts.PUT("/:id", contactController.Update)
		contacts.DELETE("/:id", contactController.Delete)
	}

	return r
}
