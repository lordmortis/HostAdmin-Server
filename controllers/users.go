package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/services"
	"github.com/lordmortis/HostAdmin-Server/viewmodels"
	"github.com/volatiletech/sqlboiler/boil"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

var (
	dbService *services.DatabaseService
)

func Users(router *gin.Engine, db *services.DatabaseService) {
	dbService = db
	router.GET("/1/users", list)
	router.POST("/1/users", create)
	router.GET("/1/user/:id", get)
}

func list(ctx *gin.Context) {
	dbModels, err := datamodels_raw.Users().All(ctx, (*dbService).GetConnection())
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
	ctx.JSON(200, viewModels)
 }

func get(ctx *gin.Context) {
	dbModel, err := datamodels.UserById(ctx, dbService, ctx.Param("id"))
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
	ctx.JSON(200, viewModel)
}

func create(ctx *gin.Context) {
	newUserJson := viewmodels.User{}

	if err := ctx.ShouldBindJSON(&newUserJson); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	if !newUserJson.Validate() {
		_ = ctx.Error(errors.New("invalid model"))
		return
	}

	dbModel := datamodels_raw.User{}
	newUserJson.ToDB(&dbModel)

	if err := dbModel.Insert(ctx, (*dbService).GetConnection(), boil.Infer()); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	if err := dbModel.Reload(ctx, (*dbService).GetConnection()); err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	newUserJson = viewmodels.User{}
	newUserJson.FromDB(&dbModel)
	ctx.JSON(200, newUserJson)
}