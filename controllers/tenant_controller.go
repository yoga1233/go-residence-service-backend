package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

// TenantController struct
type TenantController struct {
	TenantService service.TenantService
}

// Constructor untuk TenantController
func NewTenantController(tenantService service.TenantService) *TenantController {
	return &TenantController{TenantService: tenantService}
}

// GetTenants - Mendapatkan daftar tenant
func (tc *TenantController) GetTenants(c *fiber.Ctx) error {
	tenants, err := tc.TenantService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, tenants))
}

// CreateTenant - Membuat tenant baru
func (tc *TenantController) CreateTenant(c *fiber.Ctx) error {
	tenant := new(model.Tenant)
	if err := c.BodyParser(tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}
	if err := tc.TenantService.CreateTenant(tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot create tenant", fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("tenant created", fiber.StatusOK, tenant))
}

// DeleteTenant - Menghapus tenant berdasarkan ID
func (tc *TenantController) DeleteTenant(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}
	if err := tc.TenantService.DeleteTenant(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("tenant deleted", fiber.StatusOK, []string{}))
}

// UpdateTenant - Memperbarui informasi tenant
func (tc *TenantController) UpdateTenant(c *fiber.Ctx) error {
	tenant := new(model.Tenant)
	if err := c.BodyParser(tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}
	if err := tc.TenantService.UpdateTenant(tenant); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot update tenant", fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("tenant updated", fiber.StatusOK, tenant))
}

// FindTenantByQuery - Mencari tenant berdasarkan query (nama, kategori, dll.)
func (tc *TenantController) FindTenantByQuery(c *fiber.Ctx) error {
	query := c.Params("query")
	tenant, err := tc.TenantService.FindTenantByQuery(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("tenant found", fiber.StatusOK, tenant))
}
