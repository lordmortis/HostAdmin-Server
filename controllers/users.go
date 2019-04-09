package controllers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/viewmodels"
	"github.com/volatiletech/sqlboiler/boil"
	"net/http"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

func Users(router *gin.Engine) {
	userGroup := router.Group("/1/users")
	userGroup.GET("", list)
	userGroup.POST("", create)
	userGroup.GET("/:id", show)
	userGroup.PUT("/:id", update)
	userGroup.DELETE("/:id", delete)
}

func list(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModels, err := datamodels_raw.Users().All(ctx, dbCon)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	viewModels := make([]viewmodels.User, len(dbModels))
	for index := range dbModels {
		viewModel := viewmodels.User{}
		viewModel.FromDB(dbModels[index])
		viewModels[index] = viewModel
	}

	JSONOk(ctx, viewModels)
 }

func show(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
 	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
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

	viewModel := viewmodels.User{}
	viewModel.FromDB(dbModel)
	JSONOk(ctx, viewModel)
}

func create(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	newUserJson := viewmodels.User{}

	if err := ctx.ShouldBindJSON(&newUserJson); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	if !newUserJson.Validate() || len(newUserJson.NewPassword) == 0 {
		_ = ctx.Error(errors.New("invalid model"))
		return
	}

	dbModel := datamodels_raw.User{}
	newUserJson.ToDB(&dbModel)

	if err := dbModel.Insert(ctx, dbCon, boil.Infer()); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	if err := dbModel.Reload(ctx, dbCon); err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	newUserJson = viewmodels.User{}
	newUserJson.FromDB(&dbModel)
	JSONOk(ctx, newUserJson)
}

func update(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
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

	newUserJson := viewmodels.User{}
	if err := ctx.ShouldBindJSON(&newUserJson); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	newUserJson.ToDB(dbModel)

	rows, err := dbModel.Update(ctx, dbCon, boil.Infer())
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	newUserJson = viewmodels.User{}
	newUserJson.FromDB(dbModel)
	if rows == 1 {
		JSONOk(ctx, newUserJson)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
	}
}

func delete(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
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
		JSONBadRequest(ctx, gin.H{"general": [1]string{"unable to delete"}})
	}
}