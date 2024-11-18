package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"ecommerce-api/config"
)

func GenerateToken( userId string) string {

	claims := jwt.MapClaims{
		"userId":  userId,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	t, _ := token.SignedString([]byte(config.Config("JWT_SECRET_KEY")))

	return t
}
