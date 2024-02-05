package routes

import (
	"fmt"
	v1 "outsource-management/api/routes/v1"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(host string, port string) error {
	app := fiber.New()

	app.Use(cors.New())

	api := app.Group("/api")
	v1.InitApiV1(api.Group("/v1"))

	return app.Listen(fmt.Sprintf("%s:%s", host, port))
}
