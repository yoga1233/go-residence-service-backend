package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yoga1233/go-residence-service-backend/config"
)

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET_KEY", "secret")))
}
