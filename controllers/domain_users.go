package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/pkg/errors"
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
	domain := fetchDomain(ctx)
	if domain == nil {
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
	//TODO: Validate user permissions - check that the user can administer domain.
	domain := fetchDomain(ctx)
	if domain == nil {
		return
	}

	model := datasource.UserDomain{}

	if err := ctx.ShouldBindJSON(&model); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	model.Domain = domain
	if model.UserID != "" {
		tempID := datasource.UUIDFromString(model.UserID)
		if tempID != uuid.Nil {
			model.UserID = tempID.String()
		}
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

func showDomainUser(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainUser(ctx)
	if model == nil { return }

	JSONOk(ctx, model)
}

func updateDomainUser(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainUser(ctx)
	if model == nil { return }

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

func deleteDomainUser(ctx *gin.Context) {
	//TODO: Validate user permissions - check that the user can administer domain.
	model := fetchDomainUser(ctx)
	if model == nil { return }

	updated, err := model.Delete(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if updated {
		JSONOkStatusResponse(ctx)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to deleteUser"}})
	}
}

func fetchDomainUser(ctx *gin.Context) *datasource.UserDomain{
	domain := fetchDomain(ctx)
	if domain == nil { return nil }

	userID := datasource.UUIDFromString(ctx.Param("user_id"))
	if userID == uuid.Nil {
		JSONBadRequest(ctx, gin.H{"id": [1]string{"Unable to parse user ID"}})
		return nil
	}

	model, err := datasource.UserDomainsWithIDs(ctx, userID, domain.IDUuid, true, false)
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
