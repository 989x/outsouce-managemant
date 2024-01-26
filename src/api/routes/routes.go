package routes

import (
	"osm/api/controllers"
	"osm/api/handler"
	"osm/api/models"

	auth_repository "osm/api/repository/auth_repo"

	auth_service "osm/api/service/auth_service"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, mgConn *models.MongoInstance) {

	api := app.Group("/api")
	v1 := api.Group("/v1")

	authRep := auth_repository.NewAuthRepository(mgConn)
	authSrv := auth_service.NewAuthRepository(authRep)
	authHan := handler.NewAuthService(authSrv)

	auth := v1.Group("/auth")

	auth.Post("/login", authHan.Old_login)

	staff := v1.Group("/staff")

	staff.Get("/staffdashboard", controllers.GetStaffDashBoard)
	staff.Get("/staffdashboard/:project", controllers.GetStaffByJobName)
	staff.Get("/staffs", controllers.GetStaff)

}
