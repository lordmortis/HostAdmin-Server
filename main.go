package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/controllers"
	"runtime"

	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/lordmortis/HostAdmin-Server/services"
)

var (
	config *Config
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := flag.String("config", "config.json", "JSON Config file")

	flag.Parse()

	var err error

	config, err = LoadConfig(*configFile)
	if err != nil {
		fmt.Println("Unable to parse Config file")
		fmt.Println(err)
		return
	}

	err = datasource.PerformMigrations(config.Database)
	if err != nil {
		fmt.Println("Unable to perform/check migrations")
		fmt.Println(err)
		return
	}

	var dbService services.DatabaseService
	dbService, err = services.NewDatabaseService(config.Database)
	if err != nil {
		fmt.Println("Unable to setup database connection:")
		fmt.Println(err)
		return
	}


	router := gin.Default()
	controllers.Users(router, &dbService)

	router.Run(config.Server.String())
}
