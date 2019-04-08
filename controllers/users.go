package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/services"
	"github.com/lordmortis/HostAdmin-Server/viewmodels"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

var (
	dbService *services.DatabaseService
)

func Users(router *gin.Engine, db *services.DatabaseService) {
	dbService = db
	router.GET("/1/users", list)
}

func list(ctx *gin.Context) {
	dbUsers, err := datamodels_raw.Users().All(ctx, (*dbService).GetConnection())
	if err != nil {
		println(err.Error())
		_ = ctx.Error(err)
		return
	}

	jsonUsers := make([]viewmodels.User, len(dbUsers))
	for index := range dbUsers {
		jsonUsers[index] = viewmodels.UserViewModel(dbUsers[index])
	}
	ctx.JSON(200, jsonUsers)
 }
