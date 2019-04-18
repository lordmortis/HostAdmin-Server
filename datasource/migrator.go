package datasource

import (
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/lordmortis/HostAdmin-Server/config"
)

func PerformMigrations(config config.DatabaseConfig) error {
	var connString = "postgres://" + config.Username + ":" + config.Password
	connString += "@" + config.Hostname + ":" + strconv.Itoa(config.Port)
	connString += "/" + config.Database + "?sslmode=disable"
	m, err := migrate.New("file://datasource/migrations/", connString)
	if err != nil {
		return err
	}

	err = m.Up()
	_ = err
	// TODO: handle errors with migrations...

	m.Close()
	return nil
}