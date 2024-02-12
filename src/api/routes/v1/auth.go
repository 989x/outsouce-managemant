package routes

import (
	controllers "outsource-management/api/controllers/v1"
	middlewares "outsource-management/api/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RoutesAuth(v1 fiber.Router) {
	auth := v1.Group("/auth")
	auth.Post("/login", controllers.Login)
	auth.Get("/user", middlewares.RequestAuth(), controllers.Params)
}
