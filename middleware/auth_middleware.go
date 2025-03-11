package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/yoga1233/go-residence-service-backend/helper"
	"github.com/yoga1233/go-residence-service-backend/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")
	// Check if user is authenticated
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFailure("Unauthorized", fiber.StatusUnauthorized))
	}

	//hapous prefix "Bearer " dari token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Check if token is valid
	token, err := utils.VerifyJWT(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFailure("Unauthorized", fiber.StatusUnauthorized))
	}

	// Ambil klaim dari token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Ambil email dari klaim
		email, emailOk := claims["email"].(string)
		if emailOk {
			// Simpan email di context agar bisa diakses di handler berikutnya
			c.Locals("email", email)
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Ambil email dari klaim
		id, idOk := claims["id"].(string)
		if idOk {
			// Simpan email di context agar bisa diakses di handler berikutnya
			c.Locals("id", id)
		}
	}

	return c.Next()

}
