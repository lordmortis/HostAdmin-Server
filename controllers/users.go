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
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	viewModels := make([]viewmodels.User, len(dbModels))
	for index := range dbModels {
		viewModel := viewmodels.User{}
		viewModel.FromDB(dbModels[index])
		viewModels[index] = viewModel
	}
	ctx.JSON(http.StatusOK, viewModels)
 }

func show(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
 	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if dbModel == nil {
		ctx.String(404, "not found")
		return
	}

	viewModel := viewmodels.User{}
	viewModel.FromDB(dbModel)
	ctx.JSON(http.StatusOK, viewModel)
}

func create(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	newUserJson := viewmodels.User{}

	if err := ctx.ShouldBindJSON(&newUserJson); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	if !newUserJson.Validate() || len(newUserJson.NewPassword) == 0 {
		_ = ctx.Error(errors.New("invalid model"))
		return
	}

	dbModel := datamodels_raw.User{}
	newUserJson.ToDB(&dbModel)

	if err := dbModel.Insert(ctx, dbCon, boil.Infer()); err != nil {
		_ = ctx.Error(err)
		return
	}

	if err := dbModel.Reload(ctx, dbCon); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	newUserJson = viewmodels.User{}
	newUserJson.FromDB(&dbModel)
	ctx.JSON(http.StatusOK, newUserJson)
}

func update(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if dbModel == nil {
		ctx.String(http.StatusNotFound, "not found")
		return
	}

	newUserJson := viewmodels.User{}
	if err := ctx.ShouldBindJSON(&newUserJson); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	newUserJson.ToDB(dbModel)

	rows, err := dbModel.Update(ctx, dbCon, boil.Infer())
	if err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	newUserJson = viewmodels.User{}
	newUserJson.FromDB(dbModel)
	if rows == 1 {
		ctx.JSON(http.StatusOK, newUserJson)
	} else {
		ctx.JSON(http.StatusBadRequest, newUserJson)
	}
}

func delete(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	dbModel, err := datamodels.UserById(ctx, dbCon, ctx.Param("id"))
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	if dbModel == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "not found"})
		return
	}

	rows, err := dbModel.Delete(ctx, dbCon)

	if err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	if rows == 1 {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error"})
	}
}