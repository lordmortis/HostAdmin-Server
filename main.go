package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/lordmortis/HostaAdmin-Server/datasource"
	"github.com/lordmortis/HostaAdmin-Server/web/controllers"
	"runtime"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/lordmortis/HostaAdmin-Server/services"
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

	app := iris.New()

	app.Logger().SetLevel(config.Logging.Level)

	app.Use(recover.New())
	app.Use(logger.New())

	userController := mvc.New(app.Party("/users"))
	userController.Register(dbService)
	userController.Handle(new(controllers.UsersController))

	fmt.Println("Log level is: " + config.Logging.Level)
	app.Run(iris.Addr(config.Server.String()), iris.WithoutServerError(iris.ErrServerClosed))
}
