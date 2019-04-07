package services

import (
	"database/sql"

	_ "github.com/lib/pq"
	"strconv"
)

type DatabaseConfig struct {
	Hostname string
	Port int
	Database string
	Username string
	Password string
}

type DatabaseService interface {
	GetConnection() *sql.DB
}

type dbService struct {
	conn *sql.DB
}

func NewDatabaseService(config DatabaseConfig) (DatabaseService, error) {
	var connectionString = "user='" + config.Username + "' "
	connectionString += "password='" + config.Password + "' "
	connectionString += "host='" + config.Hostname + "' "
	connectionString += "port=" + strconv.Itoa(config.Port) + " "
	connectionString += "dbname='" + config.Database + "' "
	connectionString += "sslmode='disable'"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	var service = dbService{conn:db}
	return &service, nil
}

func (svc *dbService) GetConnection() *sql.DB {
	return svc.conn
}
