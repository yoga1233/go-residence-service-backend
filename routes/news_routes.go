package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/controllers"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func NewsRoutes(app *fiber.App) {
	newsRepo := repositories.NewNewsRepository(config.DB)
	newsService := service.NewNewsService(newsRepo)
	newsController := controllers.NewNewsController(newsService)

	news := app.Group("/news")

	news.Post("/", middleware.AuthMiddleware, newsController.CreateNews)
	news.Get("/", middleware.AuthMiddleware, newsController.FindAll)
	news.Get("/:limit", middleware.AuthMiddleware, newsController.FindByLimit)
	news.Patch("/", middleware.AuthMiddleware, newsController.UpdateNews)
	news.Delete("/:id", middleware.AuthMiddleware, newsController.DeleteNews)
}
