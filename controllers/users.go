package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/pkg/errors"
	"net/http"
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
	//TODO: Validate user permissions

	var models []datasource.User

	models, err := datasource.UsersAll(ctx)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOk(ctx, models)
 }

func showUser(ctx *gin.Context) {
	//TODO: Validate user permissions

	model, err := datasource.UserWithIDString(ctx, ctx.Param("id"))
	if err != nil {
		if err == datasource.UUIDParseError {
			JSONBadRequest(ctx, gin.H{"id": [1]string{err.Error()}})
		} else {
			JSONInternalServerError(ctx, err)
		}
		return
	}

	if model == nil {
		JSONNotFound(ctx)
		return
	}

	JSONOk(ctx, model)
}

func createUsers(ctx *gin.Context) {
	//TODO: Validate user permissions

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
	//TODO: Validate user permissions

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
	//TODO: Validate user permissions

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