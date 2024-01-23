package v0

import "github.com/gofiber/fiber/v2"

// SetupRoutes sets up all the routes for version 0
func SetupHealthCheckRoutes(app *fiber.App) {
	// Create a group for v0
	v0 := app.Group("/api/v0")

	// Ping/Pong route for testing functionality
	v0.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
}
