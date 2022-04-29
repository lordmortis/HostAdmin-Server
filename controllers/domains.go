package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/pkg/errors"
)

func Domains(router gin.IRoutes) {
	router.GET("", listDomains)
	router.POST("", createDomain)
}

func Domain(router gin.IRoutes) {
	router.GET("", showDomain)
	router.PUT("", updateDomain)
	router.DELETE("", deleteDomain)
}

func listDomains(ctx *gin.Context) {
	//TODO: Validate user permissions

	limit, offset, err := getPaginationParams(ctx)
	if err != nil {
		fmt.Printf("Unable to get pagination: %s\n", err)
	}

	models, count, err := datasource.DomainsAll(ctx, limit, offset)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOkTable(ctx, models, count)
}

func createDomain(ctx *gin.Context) {
	//TODO: Validate user permissions
	model := datasource.Domain{}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := model.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	_, err := model.Update(ctx)
	if err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	JSONOk(ctx, model)
}

func showDomain(ctx *gin.Context) {
	//TODO: Validate user permissions
	model := fetchDomain(ctx)
	if model == nil {
		return
	}
	JSONOk(ctx, model)
}

func updateDomain(ctx *gin.Context) {
	//TODO: Validate user permissions
	model := fetchDomain(ctx)
	if model == nil {
		return
	}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := model.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	_, err := model.Update(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	JSONOk(ctx, model)
}

func deleteDomain(ctx *gin.Context) {
	//TODO: Validate user permissions
	model := fetchDomain(ctx)
	if model == nil {
		return
	}

	deleted, err := model.Delete(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if deleted {
		JSONOkStatusResponse(ctx)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to delete domain"}})
	}
}

func fetchDomain(ctx *gin.Context) *datasource.Domain {
	domainID := datasource.UUIDFromString(ctx.Param("domain_id"))
	if domainID == uuid.Nil {
		JSONBadRequest(ctx, gin.H{"domain_id": [1]string{"Unable to parse domain ID"}})
		return nil
	}

	domain, err := datasource.DomainWithID(ctx, domainID)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return nil
	}

	if domain == nil {
		JSONBadRequest(ctx, gin.H{"domain_id": [1]string{"Unable to find domain ID"}})
		return nil
	}

	return domain
}
