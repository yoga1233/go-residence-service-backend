package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/helper"
	"github.com/yoga1233/go-residence-service-backend/middleware"
	model "github.com/yoga1233/go-residence-service-backend/models"
	"github.com/yoga1233/go-residence-service-backend/repositories"
	service "github.com/yoga1233/go-residence-service-backend/services"
	"github.com/yoga1233/go-residence-service-backend/utils"
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
		type ReportReq struct {
			Title       string `json:"title" validate:"required"`
			Description string `json:"description" validate:"required"`
		}
		id := c.Locals("id").(int)
		req := new(ReportReq)

		report := new(model.Report)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		valid := utils.Validate(req)
		if valid != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(valid.Error(), fiber.StatusBadRequest))
		}

		report.UserID = uint(id)
		report.Title = req.Title
		report.Description = req.Title
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
