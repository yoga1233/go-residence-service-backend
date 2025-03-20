package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

// TenantOrderController struct
type TenantOrderController struct {
	TenantOrderService service.TenantOrderService
}

// Constructor untuk TenantOrderController
func NewTenantOrderController(tenantOrderService service.TenantOrderService) *TenantOrderController {
	return &TenantOrderController{TenantOrderService: tenantOrderService}
}

// DTO untuk membuat pesanan tenant
type CreateTenantOrderDTO struct {
	TenantID int `json:"tenant_id" validate:"required"`
}

// GetTenantOrders - Mendapatkan daftar pesanan tenant berdasarkan user ID
func (tc *TenantOrderController) GetTenantOrders(c *fiber.Ctx) error {
	id := c.Locals("id").(int)

	tenantOrders, err := tc.TenantOrderService.FindByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrders))
}

// CreateTenantOrder - Membuat pesanan tenant baru
func (tc *TenantOrderController) CreateTenantOrder(c *fiber.Ctx) error {
	var dto CreateTenantOrderDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	id := c.Locals("id").(int)

	tenantOrder := &model.TenantOrder{
		UserID:   id,
		TenantID: dto.TenantID,
	}

	if err := tc.TenantOrderService.CreateTenantOrder(tenantOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot create tenant order", fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
}

// UpdateTenantOrder - Memperbarui pesanan tenant
func (tc *TenantOrderController) UpdateTenantOrder(c *fiber.Ctx) error {
	tenantOrder := new(model.TenantOrder)
	if err := c.BodyParser(tenantOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	if err := tc.TenantOrderService.UpdateTenantOrder(tenantOrder); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot update tenant order", fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenantOrder))
}
