package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
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

	authMiddleware := middleware.Auth()
	middleware.AuthSetConfig(conf.Auth)

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = conf.Server.AllowedOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
	router.Use(dbMiddleware)
	router.Use(redisMiddleware)

	loginGroup := router.Group("/1/login")
	controllers.Login(loginGroup)

	sessionKeepalive := router.Group("/1/auth")
	sessionKeepalive.Use(authMiddleware)
	controllers.Auth(sessionKeepalive)

	userGroup := router.Group("/1/users")
	userGroup.Use(authMiddleware)
	controllers.Users(userGroup)

	domainGroup := router.Group("/1/domains")
	domainGroup.Use(authMiddleware)
	controllers.Domains(domainGroup)

	err = router.Run(conf.Server.String())
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}