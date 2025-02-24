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

func TenantOrderRoutes(app *fiber.App) {

	tenantOrderRepo := repositories.NewTenantOrderRepository(config.DB)
	tenantOrderService := service.NewTenantOrderService(tenantOrderRepo)

	app.Get("/tenantOrder/:userId", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("userId"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		tenantOrder, err := tenantOrderService.FindByUserID(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})

	app.Post("/tenantOrder", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		tenantOrder := new(model.TenantOrder)
		if err := c.BodyParser(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := tenantOrderService.CreateTenantOrder(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("cannot create tenant order", fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})

	app.Patch("/tenantOrder", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		tenantOrder := new(model.TenantOrder)
		if err := c.BodyParser(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := tenantOrderService.UpdateTenantOrder(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponseFailure("cannot update tenant order", fiber.StatusBadRequest))
		}
		return c.JSON(response.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})
}
