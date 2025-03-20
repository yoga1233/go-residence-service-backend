package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	controllers "github.com/yoga1233/go-residence-service-backend/controller"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func AuthRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	auth := app.Group("/auth")

	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)
	auth.Get("/status", middleware.AuthMiddleware, authController.Status)
}
