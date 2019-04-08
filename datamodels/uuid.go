package datamodels

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
)

func UUIDFromString(uuidString string) uuid.UUID {
	var realUUID uuid.UUID
	uuidBytes, err := base64.StdEncoding.DecodeString(uuidString)
	if err == nil {
		realUUID = uuid.FromBytesOrNil(uuidBytes)
	} else {
		realUUID = uuid.FromStringOrNil(uuidString)
	}

	return realUUID
}