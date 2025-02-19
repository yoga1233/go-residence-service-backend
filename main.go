package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/yoga1233/go-residence-service-backend/config"
	"github.com/yoga1233/go-residence-service-backend/routes"
)

func main() {
	config.ConnectDB()
	app := fiber.New()

	routes.AuthRoutes(app)

	log.Println("Server is running on http://localhost:3000")
	app.Listen(":3000")
}
