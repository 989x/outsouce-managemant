package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"dododev/api/configs"
	"dododev/api/routes"
	v0 "dododev/api/routes/v0"
	v1 "dododev/api/routes/v1"
)

func main() {
	configs.LoadEnvVariables()

	app := fiber.New()

	app.Use(logger.New())

	app.Use(recover.New())

	configs.InitMongoDB()

	routes.SetupIndexRoutes(app)
	v0.SetupHealthCheckRoutes(app)
	v1.SetupUserRoutes(app)

	port := configs.GetPort()

	err := app.Listen(":" + port)
	if err != nil {
		panic(err)
	}
}
