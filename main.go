package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/routes"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TenantRoutes(app)
	routes.NewsRoutes(app)
	routes.TenantOrderRoutes(app)

	log.Println("Server is running on http://localhost:3000")
	app.Listen("0.0.0.0:3000")
}
