package routes

import (
	controllers "outsource-management/api/controllers/v1"

	"github.com/gofiber/fiber/v2"
)

func RoutesStaff(v1 fiber.Router) {
	staff := v1.Group("/staff")
	staff.Get("/testdashboard", controllers.GetStaffDashBoardTest)
	staff.Get("/staffdashboard", controllers.GetStaffDashBoard)
	staff.Get("/staffdashboard/:project", controllers.GetStaffByJobName)
	staff.Get("/staffs", controllers.GetStaff)
	staff.Get("/staffs/:id", controllers.GetStaffView)
	staff.Get("staffjobs/:id", controllers.GetStaffJobView)
	staff.Post("/fillter", controllers.GetFillterStaff)
	staff.Put("/staffs/:id", controllers.UpdateStaff)
	staff.Get("/skills", controllers.GetAllSkill)
}
