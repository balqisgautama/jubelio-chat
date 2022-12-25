package main

import (
	"github.com/balqisgautama/jubelio-chat/config"
	"github.com/balqisgautama/jubelio-chat/config/server"
	"github.com/balqisgautama/jubelio-chat/http/router"
	"github.com/balqisgautama/jubelio-chat/seeder"
	"github.com/balqisgautama/jubelio-chat/util"
	"os"
)

func main() {
	var arguments = "development"
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	server.SetServerConfig()
	seeder.DBMigrate()
	util.InitializeLogger()

	router.ApiController(config.ApplicationConfiguration.GetServerPort())
}
