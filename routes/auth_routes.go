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

func AuthRoutes(app *fiber.App) {

	userRepo := repositories.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo)

	auth := app.Group("/auth")

	auth.Post("/register", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		err := utils.Validate(user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		if err := authService.Register(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}
		return c.JSON(helper.ApiResponseSuccess("user created", fiber.StatusOK, []string{}))
	})

	auth.Post("/login", func(c *fiber.Ctx) error {
		type userRequest struct {
			Email    string `json:"email" validate:"required"`
			Password string `json:"password" validate:"required"`
		}
		login := new(userRequest)
		if err := c.BodyParser(login); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure("invalid request", fiber.StatusBadRequest))
		}

		err := utils.Validate(login)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}

		userResponse, err := authService.Login(login.Email, login.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusUnauthorized))
		}
		return c.JSON(helper.ApiResponseSuccess("login success", fiber.StatusOK, userResponse))
	})

	app.Get("/status", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		email := c.Locals("email").(string)

		result, err := authService.Status(email)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ApiResponseFailure(err.Error(), fiber.StatusBadRequest))
		}

		return c.JSON(helper.ApiResponseSuccess("login success", fiber.StatusOK, map[string]interface{}{
			"reports":       result.Reports,
			"tenant_orders": result.TenantOrders,
		}))

	})

}
