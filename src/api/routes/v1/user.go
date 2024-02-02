package v1

import (
	"github.com/gofiber/fiber/v2"

	"dododev/api/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	userRoute := app.Group("api/v1/users")

	userRoute.Get("/", handlers.GetAllUsers)
}
