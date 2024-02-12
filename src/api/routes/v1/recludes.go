package routes

import (
	controllers "outsource-management/api/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func RoutesRecludes(v1 fiber.Router) {
	recludes := v1.Group("/recludes")

	recludes.Get("/recludes", controllers.MethodGet)
	recludes.Post("/recludes", controllers.MethodPost)
	recludes.Put("/recludes", controllers.MethodPut)
	recludes.Delete("/recludes", controllers.MethodDelete)
}
