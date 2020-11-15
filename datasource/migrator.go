package datasource

import (
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"

	"github.com/lordmortis/HostAdmin-Server/config"
	"github.com/lordmortis/HostAdmin-Server/datasource/migrationData"
)

func PerformMigrations(config config.DatabaseConfig, development bool) error {
	var connString = "postgres://" + config.Username + ":" + config.Password
	connString += "@" + config.Hostname + ":" + strconv.Itoa(config.Port)
	connString += "/" + config.Database + "?sslmode=disable"

	var m *migrate.Migrate
	var err error

	if development {
		m, err = migrate.New("file://datasource/migrations", connString)
	} else {
		s := bindata.Resource(migrationData.AssetNames(),
			func(name string) ([]byte, error) {
				return migrationData.Asset(name)
			})
		d, err := bindata.WithInstance(s)
		if err != nil {
			return err
		}
		m, err = migrate.NewWithSourceInstance("go-bindata", d, connString)
	}

	if err != nil {
		return err
	}
	
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange{
		return err
	}

	srcErr, dstErr := m.Close()

	if srcErr != nil {
		return srcErr
	}

	if dstErr != nil {
		return dstErr
	}

	return nil
}