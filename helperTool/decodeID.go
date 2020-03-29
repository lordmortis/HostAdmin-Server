package main

import (
	"encoding/base64"
	"encoding/hex"

	//	"encoding/base64"
//	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"

	//	"github.com/gofrs/uuid"
	"regexp"
)

type DecodeIDCommand struct {
}

var decodeIDCommand DecodeIDCommand

func init() {
	_, err := parser.AddCommand(
		"decodeID",
		"Decode ID",
		"Decode an id into the various formats",
		&decodeIDCommand)
	if err != nil {
		panic(err)
	}
}

func (x *DecodeIDCommand)Execute(args[]string) error {
	if len(args) < 1 {
		return errors.New("no id provided")
	}

	reg, err := regexp.Compile("[^a-zA-Z0-9+/]+")
	if err != nil {
		return err
	}

	stringID := reg.ReplaceAllString(args[0], "")

	var uuidBytes []byte
	idtype := ""

	switch len(stringID) {
	case 32:
		uuidBytes, idtype, err = decodeHex(stringID)
	case 22:
		uuidBytes, idtype, err = decodeBase64(stringID)
	case 23:
		uuidBytes, idtype, err = decodeBase64(stringID)
	case 24:
		uuidBytes, idtype, err = decodeBase64(stringID)
	default:
		err = errors.New(fmt.Sprintf("unable to decode provided ID %s", args[0]))
	}

	if err != nil {
		return err
	}

	decodedID := uuid.FromBytesOrNil(uuidBytes)
	if decodedID == uuid.Nil {
		return errors.New(fmt.Sprintf("unable to decode provided ID %s", args[0]))
	}

	fmt.Printf("%s is %s\n", stringID, idtype)
	fmt.Printf("UUID: %s\n", decodedID.String())
	fmt.Printf("Base64 Encoded: %s\n", base64.StdEncoding.EncodeToString(decodedID.Bytes()))
	fmt.Printf("hex: %s\n", hex.EncodeToString(decodedID.Bytes()))
	return nil
}

func decodeHex(hexString string) ([]byte, string, error) {
	uuidBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil, "", err
	}

	return uuidBytes, "Hex String", err
}

func decodeBase64(base64String string) ([]byte, string, error) {
	uuidBytes, err := base64.RawStdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, "", err
	}

	return uuidBytes, "Base64 String", err
}