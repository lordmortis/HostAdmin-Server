package main

import (
	"database/sql"
	"fmt"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
	"github.com/lordmortis/HostAdmin-Server/config"
	"github.com/pkg/errors"
	"os"
	"runtime"
)

type Options struct {
	DBRootUser string `long:"dbRootUser" description:"root username for the database in config" default:"postgres"`
	DBRootPw   string `long:"dbRootPW" description:"root password for the database in config" default:"rootpassword"`
	ConfigFile string `long:"configFile" description:"path to config.json file" default:"../config.json"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)
var configFile *config.Config

func dbConnect(username string, password string, database string) (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"user='%s' password='%s' database='%s' host='%s' port=%d sslmode='disable'",
		username,
		password,
		database,
		configFile.Database.Hostname,
		configFile.Database.Port,
	)

	return sql.Open("postgres", connectionString)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser.CommandHandler = func(command flags.Commander, args []string) error {
		var err error

		configFile, err = config.Load(&options.ConfigFile)
		if err != nil {
			return errors.Wrap(err, "Unable to parse config file")
		}

		return command.Execute(args)
	}

	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
