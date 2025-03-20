package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

// NewsController struct
type NewsController struct {
	NewsService service.NewsService
}

// Constructor untuk NewsController
func NewNewsController(newsService service.NewsService) *NewsController {
	return &NewsController{NewsService: newsService}
}

// DTO untuk membuat berita baru
type CreateNewsDTO struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// DTO untuk memperbarui berita
type UpdateNewsDTO struct {
	ID      int    `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// CreateNews - Menambahkan berita baru
func (nc *NewsController) CreateNews(c *fiber.Ctx) error {
	var dto CreateNewsDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	news := &model.News{
		Title:   dto.Title,
		Content: dto.Content,
	}

	if err := nc.NewsService.CreateNews(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot create news", fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("news created", fiber.StatusOK, nil))
}

// FindAll - Mendapatkan semua berita
func (nc *NewsController) FindAll(c *fiber.Ctx) error {
	news, err := nc.NewsService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, news))
}

// FindByLimit - Mendapatkan berita dengan batasan jumlah
func (nc *NewsController) FindByLimit(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	news, err := nc.NewsService.FindByLimit(limit)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, news))
}

// UpdateNews - Memperbarui berita
func (nc *NewsController) UpdateNews(c *fiber.Ctx) error {
	var dto UpdateNewsDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	news := &model.News{

		Title:   dto.Title,
		Content: dto.Content,
	}
	news.ID = uint(dto.ID)

	if err := nc.NewsService.UpdateNews(news); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot update news", fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("news updated", fiber.StatusOK, nil))
}

// DeleteNews - Menghapus berita berdasarkan ID
func (nc *NewsController) DeleteNews(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	if err := nc.NewsService.DeleteNews(id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("cannot delete news", fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("news deleted", fiber.StatusOK, nil))
}
