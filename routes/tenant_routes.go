package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/models/response"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func TenantRoutes(app *fiber.App) {

	tenantRepo := repositories.NewTenantRepository(config.DB)
	tenantService := service.NewTenantService(tenantRepo)

	// tenant := app.Group("/auth")

	app.Get("/tenants", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		tenants, err := tenantService.FindAll()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("success", fiber.StatusOK, tenants))
	})

	app.Post("/tenant", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		tenant := new(model.Tenant)
		if err := c.BodyParser(tenant); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "invalid request",
			})
		}
		if err := tenantService.CreateTenant(tenant); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"message": "tenant created",
		})
	})

}
