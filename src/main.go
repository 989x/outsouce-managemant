package main

import (
	"fmt"
	"net"
	"os"
	"osm/api/database"
	"osm/api/routes"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/alecthomas/kingpin/v2"
)

const (
	ProgramName = "Outsource managrment"
	Version     = "0.0.1"
)

var (
	startArgs = struct {
		host *net.IP
		port *string
	}{}
)

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf(ProgramName, Version))
	a.Version(Version)
	a.HelpFlag.Short('h')

	startCommand := a.Command("start", "Start server command ...")
	startArgs.host = startCommand.Flag("host", "Set server host address").Envar("SERVER_HOST").Default("0.0.0.0").IP()
	startArgs.port = startCommand.Flag("post", "Set server listen port").Envar("SERVER_PORT").Default("5000").String()

	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCommand.FullCommand():
		fmt.Println("Hello command.")
	default:
	}

	// Old Service
	database.MgInit()
	fmt.Println("Success to connected MongoDB.")
	app := fiber.New()
	app.Use(cors.New())

	routes.Routes(app)

	app.Listen("0.0.0.0:3000")
}
