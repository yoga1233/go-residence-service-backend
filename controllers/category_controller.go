package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/helper"
	service "github.com/yoga1233/go-residence-service-backend/services"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

func (cc *CategoryController) GetAll(c *fiber.Ctx) error {
	result, err := cc.categoryService.FindAll()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}
	return c.JSON(helper.ApiResponseSuccess("success", fiber.StatusOK, result))

}
