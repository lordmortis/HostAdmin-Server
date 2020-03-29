package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/pkg/errors"
	"net/http"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

func Users(router gin.IRoutes) {
	router.GET("", listUsers)
	router.POST("", createUsers)
}

func User(router gin.IRoutes) {
	router.GET("/:id", showUser)
	router.PUT("/:id", updateUser)
	router.DELETE("/:id", deleteUser)
}

func listUsers(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	dbModels, err := datamodels_raw.Users().All(ctx, dbCon)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}
	viewModels := make([]datasource.User, len(dbModels))
	for index := range dbModels {
		viewModel := datasource.User{}
		viewModel.FromDB(dbModels[index])
		viewModels[index] = viewModel
	}

	JSONOk(ctx, viewModels)
 }

func showUser(ctx *gin.Context) {
 	user, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.UUIDParseError {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if user == nil {
		JSONNotFound(ctx)
		return
	}

	JSONOk(ctx, user)
}

func createUsers(ctx *gin.Context) {
	user := datasource.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := user.ValidateCreate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	_, err := user.Update(ctx)
	if err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	JSONOk(ctx, user)
}

func updateUser(ctx *gin.Context) {
	user, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.UUIDParseError {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{err.Error()}})
		return
	}

	modelErrors := user.ValidateUpdate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	if len(user.NewPassword) > 0 {
		if !user.ValidatePassword(user.OldPassword) {
			JSONBadRequest(ctx, gin.H{"current_password": [1]string{"not set or incorrect"}})
			return
		}
	}

	updated, err := user.Update(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		_ = ctx.Error(err)
		return
	}

	if updated {
		JSONOk(ctx, user)
	} else {
		JSONBadRequest(ctx, gin.H{"general": [1]string{"no rows updated"}})
	}
}

func deleteUser(ctx *gin.Context) {
	user, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updated, err := user.Delete(ctx)
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