package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yash-sojitra-20/address-book-backend/internal/config"
)

var jwtSecret string
func SetJwtSecret(c *config.Config){
	jwtSecret = c.JwtSecret
	// fmt.Println(jwtSecret)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
			return
		}

		tokenStr := parts[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// fmt.Println(jwtSecret)
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))

		c.Set("user_id", userID)
		c.Next()
	}
}
