package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"dododev/api/configs"
	"dododev/api/routes"
	v0 "dododev/api/routes/v0"
	v1 "dododev/api/routes/v1"
)

func main() {
	// Load environment variables from .env file
	configs.LoadEnvVariables()

	// Create a new Fiber instance
	app := fiber.New()

	// Middleware for logging incoming requests
	app.Use(logger.New())

	// Middleware for recovering from panics
	app.Use(recover.New())

	// Initialize MongoDB
	configs.InitMongoDB()

	// Routes
	routes.SetupIndexRoutes(app)
	v0.SetupHealthCheckRoutes(app)
	v1.SetupBookRoutes(app)

	// Get port from environment variable or default to 8000
	port := configs.GetPort()

	// Start the Fiber app
	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
