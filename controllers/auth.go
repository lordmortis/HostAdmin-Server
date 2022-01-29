package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datasource"
	"github.com/lordmortis/HostAdmin-Server/middleware"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Login(router gin.IRoutes) {
	router.POST("", authLogin)
}

func authLogin(ctx *gin.Context) {
	loginData := datasource.Login{}

	if err := ctx.ShouldBindJSON(&loginData); err != nil {
		JSONBadRequest(ctx, gin.H{"general": [1]string{errors.Wrap(err, "parse error").Error()}})
		return
	}

	modelErrors := loginData.Validate()
	if len(modelErrors) > 0 {
		JSONBadRequest(ctx, modelErrors)
		return
	}

	useRealBcrypt := true
	var hashedPass []byte

	dbModels, err := datasource.UsersWithUsername(ctx, &loginData.Username)
	if err != nil && err != sql.ErrNoRows {
		JSONInternalServerError(ctx, err)
		return
	}

	if err == sql.ErrNoRows || len(dbModels) != 1 {
		hashedPass = []byte("$2a$10$xJF47sKayeXieBkh7i98ce2/Ok13IPcSWqW9eozJAM.TTsPgVNQbK")
		useRealBcrypt = false
	} else {
		hashedPass = dbModels[0].EncryptedPassword
	}

	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(loginData.Password)); err != nil {
		JSONNotAuthorizedResponse(ctx)
		return
	}

	if !useRealBcrypt {
		JSONNotAuthorizedResponse(ctx)
		return
	}

	var userIDBytes = datasource.UUIDFromString(dbModels[0].ID).Bytes()
	session, err := middleware.AuthCreateSession(ctx, "User", userIDBytes)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	JSONOk(ctx, gin.H{
		"sessionID": session.Base64ID,
		"expiry":    session.Expiry.Format(time.RFC3339),
	})
}

func Auth(router gin.IRoutes) {
	router.POST("/logout", authLogout)
	router.GET("/keepalive", authKeepAlive)
}

func authLogout(ctx *gin.Context) {
	middleware.AuthDestroySession(ctx)
	JSONOkStatusResponse(ctx)
}

func authKeepAlive(ctx *gin.Context) {
	sessionID := ctx.MustGet("SessionID").(string)
	expiryTime := ctx.MustGet("ExpiryTime").(time.Time)

	JSONOk(ctx, gin.H{
		"sessionID": sessionID,
		"expiry":    expiryTime.Format(time.RFC3339),
	})
}
