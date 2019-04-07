package services

import (
	"github.com/go-xorm/xorm"
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
	GetEngine() *xorm.Engine
}

type dbService struct {
	engine *xorm.Engine
}

func NewDatabaseService(config DatabaseConfig) (DatabaseService, error) {
	var connectionString = "user='" + config.Username + "' "
	connectionString += "password='" + config.Password + "' "
	connectionString += "host='" + config.Hostname + "' "
	connectionString += "port=" + strconv.Itoa(config.Port) + " "
	connectionString += "dbname='" + config.Database + "' "
	engine, err := xorm.NewEngine("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	var service = dbService{engine:engine}
	return &service, nil
}

func (svc *dbService) GetEngine() *xorm.Engine{
	return svc.engine
}
