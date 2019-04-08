package datamodels

import (
	"encoding/base64"
	"github.com/satori/go.uuid"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

func UserUUID(user *datamodels_raw.User) uuid.UUID {
	return uuid.FromStringOrNil(user.ID)
}

func UserUUIDBase64(user *datamodels_raw.User) string {
	uuid := UserUUID(user)
	return base64.StdEncoding.EncodeToString(uuid.Bytes())
}