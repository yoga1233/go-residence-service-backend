package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/controllers"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func TenantOrderRoutes(app *fiber.App) {
	tenantOrderRepo := repositories.NewTenantOrderRepository(config.DB)
	tenantOrderService := service.NewTenantOrderService(tenantOrderRepo)
	tenantOrderController := controllers.NewTenantOrderController(tenantOrderService)

	tenantOrders := app.Group("/tenantOrder")

	tenantOrders.Get("/", middleware.AuthMiddleware, tenantOrderController.GetTenantOrders)
	tenantOrders.Post("/", middleware.AuthMiddleware, tenantOrderController.CreateTenantOrder)
	tenantOrders.Patch("/", middleware.AuthMiddleware, tenantOrderController.UpdateTenantOrder)
}
