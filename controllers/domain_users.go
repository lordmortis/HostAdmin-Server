package controllers

import (
	"github.com/gin-gonic/gin"
)

func DomainUsers(router gin.IRoutes) {
	router.GET("", listDomainUsers)
	router.POST("", createDomainUsers)
}

func DomainUser(router gin.IRoutes) {
	router.GET("", showDomainUser)
	router.PUT("", updateDomainUser)
	router.DELETE("", deleteDomainUser)
}

func listDomainUsers(ctx *gin.Context) {
//	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	JSONNotFound(ctx)
}

func createDomainUsers(ctx *gin.Context) {
//	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	JSONNotFound(ctx)
}

func showDomainUser(ctx *gin.Context) {
	//	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	JSONNotFound(ctx)
}

func updateDomainUser(ctx *gin.Context) {
	//	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	JSONNotFound(ctx)
}

func deleteDomainUser(ctx *gin.Context) {
	//	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	JSONNotFound(ctx)
}