package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/lordmortis/HostAdmin-Server/viewmodels"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"
)

func Domains(router gin.IRoutes) {
	router.GET("", listDomains)
	router.POST("", createDomain)
	router.GET("/:id", showDomain)
	router.PUT("/:id", updateDomain)
	router.DELETE("/:id", deleteDomain)
}

func listDomains(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	count, err := datamodels_raw.Domains().Count(ctx, dbCon)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	dbModels, err := datamodels_raw.Domains().All(ctx, dbCon)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	viewModels := make([]viewmodels.Domain, len(dbModels))
	for index := range dbModels {
		viewModel := viewmodels.Domain{}
		viewModel.FromDB(dbModels[index])
		viewModels[index] = viewModel
	}

	JSONOkTable(ctx, viewModels, count)
}

func createDomain(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	newJson := viewmodels.Domain{}

	if err := ctx.ShouldBindJSON(&newJson); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := newJson.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	dbModel := datamodels_raw.Domain{}
	newJson.ToDB(&dbModel)

	if err := dbModel.Insert(ctx, dbCon, boil.Infer()); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	if err := dbModel.Reload(ctx, dbCon); err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	newJson = viewmodels.Domain{}
	newJson.FromDB(&dbModel)
	JSONOk(ctx, newJson)
}

func showDomain(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.DomainById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		if err == datamodels.UUIDParseError {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if dbModel == nil {
		JSONNotFound(ctx)
		return
	}

	viewModel := viewmodels.Domain{}
	viewModel.FromDB(dbModel)
	JSONOk(ctx, viewModel)
}

func updateDomain(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.DomainById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		if err == datamodels.UUIDParseError {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if dbModel == nil {
		JSONNotFound(ctx)
		return
	}

	updateJson := viewmodels.Domain{}

	if err := ctx.ShouldBindJSON(&updateJson); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := updateJson.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	updateJson.ToDB(dbModel)

	rows, err := dbModel.Update(ctx, dbCon, boil.Infer())
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	updateJson = viewmodels.Domain{}
	updateJson.FromDB(dbModel)
	if rows == 1 {
		JSONOk(ctx, updateJson)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
	}
}

func deleteDomain(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.DomainById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if dbModel == nil {
		JSONNotFound(ctx)
		return
	}

	rows, err := dbModel.Delete(ctx, dbCon)

	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	if rows == 1 {
		JSONOkStatusResponse(ctx)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to delete domain"}})
	}
}