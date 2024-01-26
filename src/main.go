package main

import (
	"fmt"
	"osm/api/database"
	"osm/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	mgConn, err := database.MongoInit()
	if err != nil {
		fmt.Println("Can't connetc MongoDB.")
	}

	fmt.Println("Success to connected MongoDB.")

	database.MgInit()

	app := fiber.New()
	app.Use(cors.New())

	routes.Routes(app, mgConn)

	app.Listen("0.0.0.0:3000")
}
