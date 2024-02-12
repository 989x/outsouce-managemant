package routes

import (
	controllers "outsource-management/api/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func RoutesSkill(v1 fiber.Router) {
	skill := v1.Group("/skill")
	skill.Get("/skills", controllers.GetAllSkill)
}
