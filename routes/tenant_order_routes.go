package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/helper"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func TenantOrderRoutes(app *fiber.App) {

	tenantOrderRepo := repositories.NewTenantOrderRepository(config.DB)
	tenantOrderService := service.NewTenantOrderService(tenantOrderRepo)

	app.Get("/tenantOrder", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		id := c.Locals("id").(int)

		tenantOrder, err := tenantOrderService.FindByUserID(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})

	app.Post("/tenantOrder", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		type Request struct {
			TenantId int `json:"tenant_id" gorm:"not null"`
		}
		tenantOrder := new(model.TenantOrder)
		tenantReq := new(Request)

		id := c.Locals("id").(int)
		if err := c.BodyParser(tenantReq); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		tenantOrder.UserID = id
		tenantOrder.TenantID = tenantReq.TenantId

		if err := tenantOrderService.CreateTenantOrder(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot create tenant order", fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})

	app.Patch("/tenantOrder", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		tenantOrder := new(model.TenantOrder)
		if err := c.BodyParser(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		if err := tenantOrderService.UpdateTenantOrder(tenantOrder); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot update tenant order", fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
	})
}
