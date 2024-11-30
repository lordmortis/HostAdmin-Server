package main

import (
	"flag"
	"fmt"
	"github.com/lordmortis/HostAdmin-Server/templateData"
	"html/template"
	"os"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/lordmortis/HostAdmin-Server/config"
	"github.com/lordmortis/HostAdmin-Server/controllers"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/lordmortis/HostAdmin-Server/middleware"
)

func setupTemplates(router *gin.Engine, development bool) {
	if development {
		router.LoadHTMLGlob("templates/*")
		return
	}

	t := template.New("")

	for _, name := range templateData.AssetNames() {
		file, err := templateData.AssetInfo(name)
		if err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("Unable to get info on template file '%s' - %s\n", name, err))
			continue
		}

		if file.IsDir() {
			continue
		}

		data, err := templateData.AssetString(name)
		if err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("Unable to read template file '%s' - %s\n", name, err))
			continue
		}
		t, err = t.New(name).Parse(data)
		if err != nil {
			_, _ = os.Stderr.WriteString(fmt.Sprintf("Unable to parse template file '%s' into template - %s\n", name, err))
			continue
		}
	}

	router.SetHTMLTemplate(t)
}

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

	err = datasource.PerformMigrations(conf.Database, conf.Development)
	if err != nil {
		fmt.Println("Unable to perform/check migrations")
		fmt.Println(err)
		return
	}

	var dbMiddleware gin.HandlerFunc
	dbMiddleware, err = middleware.Database(conf.Database, conf.Development)
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

	if !conf.Development {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = conf.Server.AllowedOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}
	router.Use(cors.New(corsConfig))
	router.Use(dbMiddleware)
	router.Use(redisMiddleware)
	setupTemplates(router, conf.Development)

	loginGroup := router.Group("/1/login")
	controllers.Login(loginGroup)

	sessionKeepalive := router.Group("/1/auth")
	sessionKeepalive.Use(authMiddleware)
	controllers.Auth(sessionKeepalive)

	usersGroup := router.Group("/1/users")
	usersGroup.Use(authMiddleware)
	controllers.Users(usersGroup)

	userGroup := router.Group("/1/user/:user_id")
	userGroup.Use(authMiddleware)
	controllers.User(userGroup)

	domainsGroup := router.Group("/1/domains")
	domainsGroup.Use(authMiddleware)
	controllers.Domains(domainsGroup)

	domainGroup := router.Group("/1/domain/:domain_id")
	domainGroup.Use(authMiddleware)
	controllers.Domain(domainGroup)

	domainUsersGroup := router.Group("/1/domain/:domain_id/users")
	domainUsersGroup.Use(authMiddleware)
	controllers.DomainUsers(domainUsersGroup)

	domainUserGroup := router.Group("/1/domain/:domain_id/user/:user_id")
	domainUserGroup.Use(authMiddleware)
	controllers.DomainUser(domainUserGroup)

	domainEmailUsersGroup := router.Group("/1/domain/:domain_id/emailUsers")
	domainEmailUsersGroup.Use(authMiddleware)
	controllers.DomainEmailUsers(domainEmailUsersGroup)

	domainEmailUserGroup := router.Group("/1/domain/:domain_id/emailUser/:base_address")
	domainEmailUserGroup.Use(authMiddleware)
	controllers.DomainEmailUser(domainEmailUserGroup)

	domainEmailSelfAdminGroup := router.Group("/1/domain/:domain_id/emailUser/:base_address")
	controllers.DomainEmailUserSelfAdmin(domainEmailSelfAdminGroup)

	domainEmailAliasesGroup := router.Group("/1/domain/:domain_id/emailAliases")
	domainEmailAliasesGroup.Use(authMiddleware)
	controllers.DomainEmailAliases(domainEmailAliasesGroup)

	domainEmailAliasGroup := router.Group("/1/domain/:domain_id/emailAlias/:address")
	domainEmailAliasGroup.Use(authMiddleware)
	controllers.DomainEmailAlias(domainEmailAliasGroup)

	err = router.Run(conf.Server.String())
	if err != nil {
		fmt.Println("Unable to start server")
		fmt.Println(err)
		return
	}
}
