package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yash-sojitra-20/address-book-backend/internal/utils"
	"go.uber.org/zap"
)

func GlobalRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Logger.Error("panic recovered", zap.Any("error", err))

				utils.Error(c, http.StatusInternalServerError, "internal server error")
				// c.JSON(http.StatusInternalServerError, gin.H{
				// 	"success": false,
				// 	"error":   "internal server error",
				// })
				c.Abort()
			}
		}()
		c.Next()
	}
}
