package datasource

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/gofrs/uuid"
	"regexp"
)

var (
	UUIDParseError = errors.New("unable to parse UUID")
	stringFilter *regexp.Regexp
)

func init() {
	var err error
	stringFilter, err = regexp.Compile("[^a-zA-Z0-9+/]+")
	if err != nil {
		panic(err)
	}
}

func UUIDStringToUUID(uuidString string) uuid.UUID {
	return uuid.FromStringOrNil(uuidString)
}

func UUIDToBase64(uuidValue uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(uuidValue.Bytes())
}

func UUIDToBase64Ptr(uuidValue uuid.UUID) *string {
	value := base64.RawURLEncoding.EncodeToString(uuidValue.Bytes())
	return &value
}

func UUIDStringToBase64(uuidString string) string {
	uuidVal := uuid.FromStringOrNil(uuidString)
	return base64.RawURLEncoding.EncodeToString(uuidVal.Bytes())
}

func UUIDFromString(uuidString string) uuid.UUID {
	var realUUID uuid.UUID
	var err error
	switch len(uuidString) {
	case 22:
		realUUID, err = decodeBase64(uuidString)
	case 23:
		realUUID, err = decodeBase64(uuidString)
	case 24:
		realUUID, err = decodeBase64(uuidString)
	case 32:
		realUUID, err = decodeHexString(uuidString)
	case 36:
		realUUID, err = uuid.FromString(uuidString)
	default:
		var regexString = stringFilter.ReplaceAllString(uuidString, "")
		var length = len(regexString)
		if length == 22 {
			realUUID, err = decodeBase64(regexString)
		} else if length == 32 {
			realUUID, err = decodeHexString(regexString)
		} else {
			return uuid.Nil
		}
	}

	if err != nil {
		return uuid.Nil
	}

	return realUUID
}

func decodeBase64(base64String string) (uuid.UUID, error) {
	var encoder *base64.Encoding
	if len(base64String) == 22 {
		encoder = base64.RawURLEncoding
	} else {
		encoder = base64.URLEncoding
	}

	uuidBytes, err := encoder.DecodeString(base64String)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(uuidBytes)
}

func decodeHexString(hexString string) (uuid.UUID, error) {
	uuidBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(uuidBytes)
}