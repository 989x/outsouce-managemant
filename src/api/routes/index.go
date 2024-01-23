package routes

import "github.com/gofiber/fiber/v2"

// SetupIndexRoutes sets up the route for the root path
func SetupIndexRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hi")
	})

	// Routes for health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Ping/Pong route for testing functionality
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
}
