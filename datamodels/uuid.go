package datamodels

import (
	"encoding/base64"
	"errors"
	"github.com/satori/go.uuid"
)

var UUIDParseError = errors.New("unable to parse UUID")

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