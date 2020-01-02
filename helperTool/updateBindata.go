package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kevinburke/go-bindata"
)

type UpdateBindata struct {
	Directory string `long:"directory" description:"source code root directory" default:"../"`
}

var uBindataCommand UpdateBindata

func init() {
	_, err := parser.AddCommand(
		"updateBindata",
		"Create/Update Bindata",
		"Create/Update the Bindata (migrations)",
		&uBindataCommand)

	if err != nil {
		panic(err)
	}
}

func (x *UpdateBindata)Execute(args[]string) error {
	errString := ""
	migrationDirectory := ""
	binDataDirectory := ""

	if len(x.Directory) == 0 {
		errString += "Need the project's root directory\n"
	} else {
		migrationDirectory = filepath.Join(x.Directory, "datasource", "migrations")
		stat, err := os.Stat(migrationDirectory)
		if err != nil {
			errString += fmt.Sprintf("migration directory '%s' doesn't exist.", migrationDirectory)
		} else if !stat.IsDir() {
			errString += fmt.Sprintf("migration directory '%s' isn't a directory.", migrationDirectory)
		}

		binDataDirectory = filepath.Join(x.Directory, "datasource", "migrationData")
		stat, err = os.Stat(binDataDirectory)
		if err != nil {
			errString += fmt.Sprintf("migration bindata directory '%s' doesn't exist.", binDataDirectory)
		} else if !stat.IsDir() {
			errString += fmt.Sprintf("migration bindata directory '%s' isn't a directory.", binDataDirectory)
		}
	}

	if len(errString) > 0 {
		return errors.New(errString)
	}

	config := bindata.Config{
		Package: "migrationData",
		Input: []bindata.InputConfig{bindata.InputConfig{Path: migrationDirectory, Recursive: false}},
		Output: filepath.Join(binDataDirectory, "main.go"),
		Prefix: migrationDirectory,
	}

	return bindata.Translate(&config)
}