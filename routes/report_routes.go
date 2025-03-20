package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/controllers"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

func ReportRoutes(app *fiber.App) {
	reportRepo := repositories.NewReportRepository(config.DB)
	reportService := service.NewReportService(reportRepo)
	reportController := controllers.NewReportController(reportService)

	reports := app.Group("/reports")

	reports.Get("/", middleware.AuthMiddleware, reportController.GetReports)
	reports.Post("/", middleware.AuthMiddleware, reportController.CreateReport)
	reports.Patch("/", middleware.AuthMiddleware, reportController.UpdateReport)
	reports.Delete("/:id", middleware.AuthMiddleware, reportController.DeleteReport)
}
