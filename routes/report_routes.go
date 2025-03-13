package routes

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		title := c.FormValue("title")
		desc := c.FormValue("description")
		file, err := c.FormFile("image")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("failed to uploud image", fiber.StatusBadRequest))
		}

		if title == "" || desc == "" {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("Title and description cannot be empty", fiber.StatusBadRequest))
		}

		uniqueId := uuid.New()

		filename := strings.Replace(uniqueId.String(), "-", "", -1)

		fileExt := strings.Split(file.Filename, ".")[1]

		image := fmt.Sprintf("%s.%s", filename, fileExt)

		err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

		if err != nil {

			return c.Status(fiber.ErrBadGateway.Code).JSON(helper.ApiResponseFailure("failed to save image", fiber.ErrBadGateway.Code))
		}

		imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)

		id := c.Locals("id").(int)

		report := new(model.Report)

		report.UserID = uint(id)
		report.Title = title
		report.Description = desc
		report.ImageUrl = imageUrl

		error := reportService.CreateReport(report)
		if error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(error.Error(), fiber.StatusBadRequest))
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
