package datamodels

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/null"
	"golang.org/x/crypto/bcrypt"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

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

func UserValidatePassword(user *datamodels_raw.User, password string) bool {
	if err := bcrypt.CompareHashAndPassword(user.EncryptedPassword.Bytes, []byte(password)); err != nil {
		return false
	}
	return true
}