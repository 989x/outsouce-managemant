package routes

import (
	"osm/api/handler"
	"osm/api/middleware"
	"osm/api/models"

	auth_repository "osm/api/repository/auth_repo"
	staff_repository "osm/api/repository/staff_repo"
	user_repository "osm/api/repository/user_repo"

	auth_service "osm/api/service/auth_service"
	staff_service "osm/api/service/staff_service"
	user_service "osm/api/service/user_service"

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

	userRep := user_repository.NewUserRepository(mgConn)
	userSrv := user_service.NewUserService(userRep)
	userHan := handler.NewUserHandler(userSrv)

	user := v1.Group("/user")

	user.Get("/getall", userHan.GetAll)
	user.Get("/getparamsuser", middleware.RequestAuth(), userHan.ParamsUser)

	staffRep := staff_repository.NewStaffRepoSitory(mgConn)
	staffSrv := staff_service.NewStaffService(staffRep)
	staffHan := handler.NewStaffService(staffSrv)

	staff := v1.Group("/staff")

	staff.Get("/dashboard", staffHan.GetDashboard)
	staff.Get("/getall", staffHan.ListStaffs)
	staff.Get("/getby/:id", staffHan.ReadStaff)
	staff.Get("/testGet", func(c *fiber.Ctx) error {
		result, err := staffRep.GetAllStaffJobs()
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"errors": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data": result,
		})
	})

}
