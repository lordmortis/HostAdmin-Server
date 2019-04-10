package controllers

import (
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/lordmortis/HostAdmin-Server/middleware"
	"github.com/lordmortis/HostAdmin-Server/viewmodels"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strings"
	"time"
)

func Login(router gin.IRoutes) {
	router.POST("", login)
}

func login(ctx *gin.Context) {
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	redisCon := ctx.MustGet("redisConnection").(*redis.Client)
	redisPrefix := ctx.MustGet("redisPrefix").(string)

	loginData := viewmodels.Login{}

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

	dbModels, err := datamodels_raw.Users(qm.Where("username = ?", loginData.Username)).All(ctx, dbCon)
	if err != nil && err != sql.ErrNoRows {
		JSONInternalServerError(ctx, err)
		return
	}

	if err == sql.ErrNoRows || len(dbModels) != 1 {
		hashedPass = []byte("$2a$10$xJF47sKayeXieBkh7i98ce2/Ok13IPcSWqW9eozJAM.TTsPgVNQbK")
		useRealBcrypt = false
	} else {
		hashedPass = dbModels[0].EncryptedPassword.Bytes
	}

	if err := bcrypt.CompareHashAndPassword(hashedPass, []byte(loginData.Password)); err != nil {
		JSONNotAuthorizedResponse(ctx)
		return
	}

	if !useRealBcrypt {
		JSONNotAuthorizedResponse(ctx)
		return
	}

	var userIDBytes = datamodels.UUIDFromString(dbModels[0].ID).Bytes()
	session, err := middleware.AuthCreateSession(redisCon, redisPrefix,"User", userIDBytes)
	if err != nil {
		JSONInternalServerError(ctx, err)
		return
	}

	maxAge := int(math.Round(session.Expiry.Sub(time.Now()).Seconds()))
	domain := strings.Split(ctx.Request.Host, ":")[0]
	ctx.SetCookie("sessionID", session.Base64ID, maxAge, "/", domain, true, true)

	JSONOk(ctx, gin.H{
		"sessionID": base64.StdEncoding.EncodeToString(session.ID.Bytes()),
		"expiry": session.Expiry.Format(time.RFC3339),
	})
}