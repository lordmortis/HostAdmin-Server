package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserCommand struct {
	Username string `long:"username" description:"Username for new user"`
	Email string `long:"email" description:"Email address for new user"`
	Password string `long:"password" description:"Password for new user"`
	SuperAdmin bool `long:"superAdmin" description:"If the user is a super administrator"`
}

var cuserCommand CreateUserCommand

func init() {
	_, err := parser.AddCommand(
		"createUser",
		"Create User",
		"Create user in database",
		&cuserCommand)
	if err != nil {
		panic(err)
	}
}

func (x *CreateUserCommand)Execute(args[]string) error {
	dbCon, err := dbConnect(configFile.Database.Username, configFile.Database.Password, configFile.Database.Database)

	if err != nil {
		return errors.Wrap(err, "Unable to connect to database")
	}

	errString := ""

	if len(x.Username) == 0 {
		errString += "Need a username\n"
	}

	if len(x.Email) == 0 {
		errString += "Need an email\n"
	}

	if len(x.Password) == 0 {
		errString += "Need a password\n"
	}

	if len(errString) > 0 {
		return errors.New(errString)
	}

	auuid, err := uuid.NewV4()
	if err != nil {
		return errors.Wrap(err, "Unable to generate new UUID")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(x.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "Unable to generate encrypted password")
	}

	sqlCmd := fmt.Sprintf(
		"INSERT INTO users" +
		"(id, username, email, encrypted_password, super_admin, created_at, updated_at)" +
		" VALUES " +
		"('%s', '%s', '%s', '%s', %t, now(), now()) " +
		"ON CONFLICT (username) DO " +
		"UPDATE SET " +
			"email = '%s', " +
			"encrypted_password = '%s', " +
			"super_admin = %t, " +
			"updated_at = now() ",
		auuid.String(),
		x.Username,
		x.Email,
		hashedPassword,
		x.SuperAdmin,
		x.Email,
		hashedPassword,
		x.SuperAdmin,
		)

	_, err = dbCon.Exec(sqlCmd)

	if err != nil {
		return errors.Wrap(err, "Unable to update database")
	}

	return nil
}