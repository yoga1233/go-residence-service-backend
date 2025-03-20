package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/helper"
	model "github.com/yoga1233/go-residence-service-backend/models"
	service "github.com/yoga1233/go-residence-service-backend/services"
	"github.com/yoga1233/go-residence-service-backend/utils"
)

// AuthController struct
type AuthController struct {
	AuthService service.AuthService
}

// Constructor untuk AuthController
func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

// DTO untuk registrasi
type RegisterDTO struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// DTO untuk login
type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// Register user baru
func (ac *AuthController) Register(c *fiber.Ctx) error {
	var dto RegisterDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	// Validasi input
	if err := utils.Validate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	// Hash password sebelum disimpan
	hashedPassword, err := utils.HashPassword(dto.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helper.ApiResponseFailure("failed to hash password", fiber.StatusInternalServerError))
	}

	user := &model.User{
		Username: dto.Username,
		Email:    dto.Email,
		Password: hashedPassword,
	}

	if err := ac.AuthService.Register(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.Status(fiber.StatusCreated).JSON(helper.ApiResponseSuccess("user created", fiber.StatusCreated, fiber.Map{}))
}

// Login user
func (ac *AuthController) Login(c *fiber.Ctx) error {
	var dto LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
	}

	if err := utils.Validate(dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	userResponse, err := ac.AuthService.Login(dto.Email, dto.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusUnauthorized))
	}

	return c.JSON(helper.ApiResponseSuccess("login success", fiber.StatusOK, userResponse))
}

// Status user
func (ac *AuthController) Status(c *fiber.Ctx) error {
	email, ok := c.Locals("email").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFailure("unauthorized", fiber.StatusUnauthorized))
	}

	result, err := ac.AuthService.Status(email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
	}

	return c.JSON(helper.ApiResponseSuccess("status retrieved successfully", fiber.StatusOK, fiber.Map{
		"reports":       result.Reports,
		"tenant_orders": result.TenantOrders,
	}))
}
