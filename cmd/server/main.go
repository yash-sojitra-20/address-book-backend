package main

import (
	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/yash-sojitra-20/address-book-backend/internal/db"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
	"github.com/yash-sojitra-20/address-book-backend/internal/router"
)

func main() {
	// Load config
	cfg := config.Load()

	// Init logger
	middleware.InitLogger()
	defer middleware.Logger.Sync()

	// Connect DB
	db.Connect(cfg)

	// Setup router
	r := router.Setup()

	// Start server
	r.Run(":" + cfg.ServerPort)
}
