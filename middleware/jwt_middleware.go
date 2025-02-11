package middleware

import (
	"livechat-support/config"
	"livechat-support/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, _ := token.SignedString([]byte(config.JwtSecret))
	return tokenString
}
