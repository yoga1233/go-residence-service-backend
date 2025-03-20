package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/controllers"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func TenantRoutes(app *fiber.App) {
	tenantRepo := repositories.NewTenantRepository(config.DB)
	tenantService := service.NewTenantService(tenantRepo)
	tenantController := controllers.NewTenantController(tenantService)

	tenants := app.Group("/tenant")

	tenants.Get("/", middleware.AuthMiddleware, tenantController.GetTenants)
	tenants.Post("/", middleware.AuthMiddleware, tenantController.CreateTenant)
	tenants.Patch("/", middleware.AuthMiddleware, tenantController.UpdateTenant)
	tenants.Delete("/:id", middleware.AuthMiddleware, tenantController.DeleteTenant)
	tenants.Get("/:query", middleware.AuthMiddleware, tenantController.FindTenantByQuery)
}
