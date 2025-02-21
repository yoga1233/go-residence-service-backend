package routes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/models/response"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func NewsRoutes(app *fiber.App) {

	newsRepo := repositories.NewNewsRepository(config.DB)
	newsService := service.NewNewsService(newsRepo)
	// Auth routes
	app.Post("/news", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		news := new(model.News)
		if err := c.BodyParser(news); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := newsService.CreateNews(news); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("cannot create news", fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("news created", fiber.StatusOK, []string{}))

	})

	app.Get("/news", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		news, err := newsService.FindAll()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("success", fiber.StatusOK, news))
	})

	app.Patch("/news", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		news := new(model.News)
		if err := c.BodyParser(news); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := newsService.UpdateNews(news); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("cannot update news", fiber.StatusBadRequest))
		}

		return c.JSON(response.ApiResponseSuccess("news updated", fiber.StatusOK, []string{}))

	})

	app.Delete("/news/:id", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := newsService.DeleteNews(id); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("cannot delete news", fiber.StatusBadRequest))
		}

		return c.JSON(response.ApiResponseSuccess("news deleted", fiber.StatusOK, []string{}))
	})

}
