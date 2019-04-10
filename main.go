package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime"

	"github.com/lordmortis/HostAdmin-Server/config"
	"github.com/lordmortis/HostAdmin-Server/controllers"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/lordmortis/HostAdmin-Server/middleware"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	configFile := flag.String("config", "config.json", "JSON Config file")

	flag.Parse()

	conf, err := config.Load(configFile)
	if err != nil {
		fmt.Println("Unable to parse Config file")
		fmt.Println(err)
		return
	}

	err = datasource.PerformMigrations(conf.Database)
	if err != nil {
		fmt.Println("Unable to perform/check migrations")
		fmt.Println(err)
		return
	}

	var dbMiddleware gin.HandlerFunc
	dbMiddleware, err = middleware.Database(conf.Database)
	if err != nil {
		fmt.Println("Unable to setup database connection:")
		fmt.Println(err)
		return
	}

	var redisMiddleware gin.HandlerFunc
	redisMiddleware, err = middleware.Redis(conf.Redis)
	if err != nil {
		fmt.Println("Unable to connect to redis:")
		fmt.Println(err)
		return
	}

	router := gin.Default()
	router.Use(dbMiddleware)
	router.Use(redisMiddleware)
	controllers.Users(router)

	err = router.Run(conf.Server.String())
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}
