package routes

import (
	controllers "outsource-management/api/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func RoutesReclude(v1 fiber.Router) {
	recludes := v1.Group("/recludes")

	recludes.Post("/recludes", controllers.MethodPost)
	recludes.Put("/recludes", controllers.MethodPut)
	recludes.Delete("/recludes", controllers.MethodDelete)

	recludes.Get("/monthly-attendance", controllers.GetMonthlyEmployeeAttendance)
}
