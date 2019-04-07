package datasource

import (
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/lordmortis/HostaAdmin-Server/services"
)

func PerformMigrations(config services.DatabaseConfig) error {
	var connString = "postgres://" + config.Username + ":" + config.Password
	connString += "@" + config.Hostname + ":" + strconv.Itoa(config.Port)
	connString += "/" + config.Database + "?sslmode=disable"
	println(connString)
	m, err := migrate.New("file://datasource/migrations/", connString)
	if err != nil {
		return err
	}
	m.Up()

	return nil
}