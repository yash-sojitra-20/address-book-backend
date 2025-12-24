package utils

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(code, gin.H{
		"success": false,
		"error":   msg,
	})
}

func Success(c *gin.Context, code int, data interface{}) {
	c.JSON(code, gin.H{
		"success": true,
		"data":    data,
	})
}
