package datamodels

import (
	"database/sql"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"github.com/volatiletech/null"
	"golang.org/x/crypto/bcrypt"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

func UserUUID(user *datamodels_raw.User) uuid.UUID {
	return uuid.FromStringOrNil(user.ID)
}

func UserUUIDBase64(user *datamodels_raw.User) string {
	uuid := UserUUID(user)
	return base64.StdEncoding.EncodeToString(uuid.Bytes())
}

func UserById(ctx *gin.Context, dbCon *sql.DB, stringID string) (*datamodels_raw.User, error) {
	userID := UUIDFromString(stringID)

	if userID == uuid.Nil {
		return nil, UUIDParseError
	}

	user, err := datamodels_raw.FindUser(ctx,dbCon, userID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

func UserSetPassword(user *datamodels_raw.User, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.EncryptedPassword = null.BytesFrom(hashedPassword)
	return nil
}