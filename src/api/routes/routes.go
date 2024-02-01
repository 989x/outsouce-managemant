package routes

import (
	"osm/api/controllers"
	"osm/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")

	auth.Post("/login", controllers.Login)
	auth.Get("/users", middleware.RequestAuth(), controllers.Params)

	staff := v1.Group("/staff")

	staff.Get("/staffdashboard", controllers.GetStaffDashBoard)
	staff.Get("/staffdashboard/:project", controllers.GetStaffByJobName)
	staff.Get("/staffs", controllers.GetStaff)
	staff.Get("/staffs/:id", controllers.GetStaffView)
	staff.Get("staffjobs/:id", controllers.GetStaffJobView)
	staff.Post("/fillter", controllers.GetFillterStaff)
	staff.Put("/staffs/:id", controllers.UpdateStaff)

}
