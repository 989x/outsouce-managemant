package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"outsource-management/api/configs"
	command "outsource-management/api/routes"

	"github.com/alecthomas/kingpin/v2"
)

const (
	ProgramName = "Outsource management"
	Version     = "0.0.1"
)

var (
	startArgs = struct {
		host *net.IP
		port *string
	}{}
)

func main() {

	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s %s", ProgramName, Version))
	a.Version(Version)
	a.HelpFlag.Short('h')

	startCommand := a.Command("start", "Start server command.")
	startArgs.host = startCommand.Flag("host", "Set server host address.").Envar("SERVER_HOST").Default("0.0.0.0").IP()
	startArgs.port = startCommand.Flag("post", "Set server listen port").Envar("SERVER_PORT").Default("3000").String()

	//Init Database
	configs.InitMongoDb()

	//Ruc case service
	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCommand.FullCommand():
		if err := command.Start(startArgs.host.String(), *startArgs.port); err != nil {
			panic(err.Error())
		}
	default:
		fmt.Println("Command not found.")
	}
}
