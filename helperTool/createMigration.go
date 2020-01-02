package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type CreateMigrationCommand struct {
	Name string `long:"name" description:"Migration name"`
	Directory string `long:"directory" description:"source code root directory" default:"../"`
}

var cmigrationCommand CreateMigrationCommand

func init() {
	_, err := parser.AddCommand(
		"createMigration",
		"Create Migration",
		"Create a new migration",
		&cmigrationCommand)

	if err != nil {
		panic(err)
	}
}

func (x *CreateMigrationCommand)Execute(args[]string) error {
	errString := ""
	migrationDirectory := ""

	if len(x.Name) == 0 {
		errString += "Need a name for the migration\n"
	}

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
	}

	if len(errString) > 0 {
		return errors.New(errString)
	}

	var spaceRegex = regexp.MustCompile("\\s+")
	var nonWordRegex = regexp.MustCompile("\\W")

	timestamp := time.Now().UTC()
	name := spaceRegex.ReplaceAllString(x.Name, "_")
	name = nonWordRegex.ReplaceAllString(name, "-")
	name = fmt.Sprintf("%d_%s", timestamp.Unix(), name)

	up := filepath.Join(migrationDirectory, fmt.Sprintf("%s.up.sql", name))
	down := filepath.Join(migrationDirectory, fmt.Sprintf("%s.down.sql", name))

	fmt.Printf("Will create:\n\t%s\n\t%s\nProceed? ", up, down)

	if !askForConfirmation() {
		return nil
	}

	_, err := os.Create(up)
	if err == nil {
		_, err = os.Create(down)
	}

	if err != nil {
		fmt.Printf("Couldn't create file\n")
	}

	return nil
}
