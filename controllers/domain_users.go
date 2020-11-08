package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datasource"
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
	//TODO: Validate user permissions
	domainID := datasource.UUIDFromString(ctx.Param("domain_id"))
	if domainID == uuid.Nil {
		JSONBadRequest(ctx, gin.H{"id": [1]string{"Unable to parse domain ID"}})
		return
	}

	domain, err := datasource.DomainWithID(ctx, domainID)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if domain == nil {
		JSONNotFound(ctx)
		return
	}

	models, count, err := domain.Users(ctx, true)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOkTable(ctx, models, count)
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