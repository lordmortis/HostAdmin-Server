package viewmodels

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
)

func UUIDStringToBase64(uuidString string) string {
	uuidVal := uuid.FromStringOrNil(uuidString)
	return base64.RawURLEncoding.EncodeToString(uuidVal.Bytes())
}