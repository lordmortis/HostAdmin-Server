package main

import (
	"flag"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/lordmortis/hostAdminServer/controllers"
	"runtime"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
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

	app := iris.New()

	app.Logger().SetLevel(config.Logging.Level)

	app.Use(recover.New())
	app.Use(logger.New())

	mvc.New(app.Party("/users")).Handle(new(controllers.UsersController))

	fmt.Println("Log level is: " + config.Logging.Level)
	app.Run(iris.Addr(config.Server.String()), iris.WithoutServerError(iris.ErrServerClosed))
}
