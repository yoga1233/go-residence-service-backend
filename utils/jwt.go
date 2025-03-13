package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yoga1233/go-residence-service-backend/config"
)

func GenerateJWT(email string, id int) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(), // Token expires in 1 month
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.GetEnv("JWT_SECRET_KEY", "secret")))
}

func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetEnv("JWT_SECRET_KEY", "secret")), nil
	})
}
