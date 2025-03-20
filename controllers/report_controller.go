package controllers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

// ReportController struct
type ReportController struct {
	ReportService service.ReportService
}

// Constructor untuk ReportController
func NewReportController(reportService service.ReportService) *ReportController {
	return &ReportController{ReportService: reportService}
}

// DTO untuk membuat laporan
type CreateReportDTO struct {
	Title       string `form:"title" validate:"required"`
	Description string `form:"description" validate:"required"`
}

// GetReports - Mendapatkan semua laporan berdasarkan user ID
func (rc *ReportController) GetReports(c *fiber.Ctx) error {
	id := c.Locals("id").(int)

	report, err := rc.ReportService.FindByUserId(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, report))
}

// CreateReport - Membuat laporan baru
func (rc *ReportController) CreateReport(c *fiber.Ctx) error {
	var dto CreateReportDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("failed to upload image", fiber.StatusBadRequest))
	}

	uniqueId := uuid.New()
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	if err := c.SaveFile(file, fmt.Sprintf("./images/%s", image)); err != nil {
		return c.Status(fiber.ErrBadGateway.Code).JSON(helper.ApiResponseFailure("failed to save image", fiber.ErrBadGateway.Code))
	}

	imageUrl := fmt.Sprintf("http://localhost:3000/images/%s", image)
	id := c.Locals("id").(int)

	report := &model.Report{
		UserID:      uint(id),
		Title:       dto.Title,
		Description: dto.Description,
		ImageUrl:    imageUrl,
	}

	if err := rc.ReportService.CreateReport(report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, nil))
}

// UpdateReport - Memperbarui laporan
func (rc *ReportController) UpdateReport(c *fiber.Ctx) error {
	report := new(model.Report)
	if err := c.BodyParser(report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	if err := rc.ReportService.UpdateReport(report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("report updated", fiber.StatusOK, nil))
}

// DeleteReport - Menghapus laporan berdasarkan ID
func (rc *ReportController) DeleteReport(c *fiber.Ctx) error {
	reportID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid report ID", fiber.StatusBadRequest))
	}

	if err := rc.ReportService.DeleteReport(reportID); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("report deleted", fiber.StatusOK, nil))
}
