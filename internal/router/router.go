package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/middleware"
)

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(
		gin.Logger(),
		middleware.GlobalRecovery(),
	)

	// routes will be added here later

	return r
}
