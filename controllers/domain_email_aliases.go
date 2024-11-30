package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/errors"

	"github.com/lordmortis/HostAdmin-Server/datasource"
)

func DomainEmailAliases(router gin.IRoutes) {
	router.GET("", listDomainEmailAliases)
	router.POST("", createDomainEmailAlias)
}

func DomainEmailAlias(router gin.IRoutes) {
	router.GET("", showDomainEmailAlias)
	router.PUT("", updateDomainEmailAlias)
	router.DELETE("", deleteDomainEmailAlias)
}

func listDomainEmailAliases(ctx *gin.Context) {
	//TODO: Validate user permissions
	domain := fetchDomain(ctx)
	if domain == nil {
		return
	}

	models, count, err := domain.EmailAliases(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOkTable(ctx, models, count)
}

func createDomainEmailAlias(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	domain := fetchDomain(ctx)
	if domain == nil {
		return
	}

	model := datasource.DomainEmailAlias{}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Because(err, nil, "parse error").Error()}})
		return
	}

	domainID := domain.IDUuid
	existing, err := datasource.DomainEmailAliases(ctx, domainID, model.Address, false)
	if err != nil {
		fmt.Printf("Internal error: %s", err)
		JSONInternalServerError(ctx, errors.New("Internal error"))
	}

	if existing != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"a record for this domain / base address pair exists"}})
		return
	}

	model.Domain = domain

	modelErrors := model.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	_, err = model.Update(ctx)
	if err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	JSONOk(ctx, model)
}

func showDomainEmailAlias(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainEmailAlias(ctx)
	if model == nil {
		return
	}

	JSONOk(ctx, model)
}

func updateDomainEmailAlias(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainEmailAlias(ctx)
	if model == nil {
		return
	}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	modelErrors := model.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	updated, err := model.Update(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	if updated {
		JSONOk(ctx, model)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
	}
}

func deleteDomainEmailAlias(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainEmailAlias(ctx)
	if model == nil {
		return
	}

	updated, err := model.Delete(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if updated {
		JSONOkStatusResponse(ctx)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to delete email alias"}})
	}
}

func fetchDomainEmailAlias(ctx *gin.Context) *datasource.DomainEmailAlias {
	domain := fetchDomain(ctx)
	if domain == nil {
		return nil
	}

	model, err := datasource.DomainEmailAliases(ctx, domain.IDUuid, ctx.Param("address"), false)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return nil
	}

	if model == nil {
		JSONNotFound(ctx)
		return nil
	}

	return model
}
