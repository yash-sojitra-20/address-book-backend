package utils

import (
	"time"

	"github.com/yash-sojitra-20/address-book-backend/internal/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string
func SetJwtSecret(c *config.Config){
	jwtSecret = c.JwtSecret
	// fmt.Println(jwtSecret)
}

func GenerateToken(userID uint, userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"user_email": userEmail,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Println(jwtSecret)
	return token.SignedString([]byte(jwtSecret))
}