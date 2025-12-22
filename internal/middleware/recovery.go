package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GlobalRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				Logger.Error("panic recovered", zap.Any("error", err))

				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
