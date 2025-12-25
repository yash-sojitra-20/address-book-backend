package main

import (
	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/db"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
	"github.com/yash-sojitra-20/address-book-backend/internal/models"
	"github.com/yash-sojitra-20/address-book-backend/internal/router"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
)

func main() {
	// Load config
	cfg := config.Load()

	// Init logger
	middleware.InitLogger()
	defer middleware.Logger.Sync()

	// Init validator
	utils.InitValidator()

	// Set JWT Secret
	utils.SetJwtSecret(cfg)
	middleware.SetJwtSecret(cfg)

	// Set Config
	router.SetConfig(cfg)

	// Connect DB
	db.Connect(cfg)
	db.DB.AutoMigrate(&models.User{}, &models.Address{})

	// Setup router
	r := router.Setup()

	// Start server
	r.Run(":" + cfg.ServerPort)
}
