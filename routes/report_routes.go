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

func ReportRoutes(app *fiber.App) {

	reportRepo := repositories.NewReportRepository(config.DB)
	reportService := service.NewReportService(reportRepo)

	app.Get("/reports", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		id := c.Locals("id").(int)

		report, err := reportService.FindByUserId(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, report))
	})

	app.Post("/report", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		id := c.Locals("id").(int)
		report := new(model.Report)
		if err := c.BodyParser(report); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		report.UserID = uint(id)
		err := reportService.CreateReport(report)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, []string{}))
	})

	app.Patch("/report", middleware.AuthMiddleware, func(c *fiber.Ctx) error {

		report := new(model.Report)
		if err := c.BodyParser(report); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		err := reportService.UpdateReport(report)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, []string{}))
	})

	app.Delete("/report/:id", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		reportID, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid report ID", fiber.StatusBadRequest))
		}

		err = reportService.DeleteReport(reportID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, []string{}))
	})

}
