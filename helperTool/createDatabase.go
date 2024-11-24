package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type CreateDatabaseCommand struct {
}

var cdbCommand CreateDatabaseCommand

func init() {
	_, err := parser.AddCommand(
		"createDB",
		"Create Database",
		"Create database and user according to config in config file",
		&cdbCommand)
	if err != nil {
		panic(err)
	}
}

func (x *CreateDatabaseCommand) Execute(args []string) error {
	var sqlCmd string
	username := configFile.Database.Username
	password := configFile.Database.Password
	database := configFile.Database.Database

	dbCon, err := dbConnect(options.DBRootUser, options.DBRootPw, "postgres")
	if err != nil {
		return errors.Wrap(err, "Unable to connect to database")
	}

	_, err = dbCon.Exec("SELECT true")

	if err != nil {
		return errors.Wrap(err, "Unable to connect to database")
	}

	sqlCmd = fmt.Sprintf("DROP DATABASE IF EXISTS \"%s\"", database)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Wrap(err, "Unable to remove database '"+database+"'")
	}

	sqlCmd = fmt.Sprintf("DROP ROLE IF EXISTS \"%s\"", username)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Wrap(err, "Unable to remove existing user '"+username+"'")
	}

	sqlCmd = fmt.Sprintf("CREATE ROLE \"%s\" WITH PASSWORD '%s' LOGIN", username, password)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Wrap(err, "Unable to create user '"+username+"'")
	}

	sqlCmd = fmt.Sprintf("CREATE DATABASE \"%s\" OWNER \"%s\"", database, username)
	if _, err = dbCon.Exec(sqlCmd); err != nil {
		return errors.Wrap(err, "Unable to create database '"+database+"'")
	}

	return nil
}
