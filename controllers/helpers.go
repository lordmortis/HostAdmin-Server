package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSONNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"errors" : gin.H { "general": [1]string{"not found"}}})
}

func JSONBadRequest(ctx *gin.Context, errorObj gin.H) {
	ctx.JSON(http.StatusBadRequest, gin.H{"errors" : errorObj})
}

func JSONInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"errors": gin.H{ "general": [1]string{err.Error()} } })
	_ = ctx.Error(err)
}

func JSONOk(ctx *gin.Context, response interface{}) {
	ctx.JSON(http.StatusOK, response)
}

func JSONOkStatusResponse(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status" : "ok"})
}