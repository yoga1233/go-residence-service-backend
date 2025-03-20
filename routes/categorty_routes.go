package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/controllers"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func CategoryRoutes(app *fiber.App) {
	categoryRepo := repositories.NewCategoryRepository(config.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	app.Get("/categpries", middleware.AuthMiddleware, categoryController.GetAll)
}
