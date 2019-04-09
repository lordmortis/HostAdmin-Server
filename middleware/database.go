package middleware

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"

	"github.com/lordmortis/HostAdmin-Server/config"
)

func Database(config config.DatabaseConfig) (gin.HandlerFunc, error) {
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

	handler := func(ctx *gin.Context) {
		ctx.Set("databaseConnection", db)
	}
	return handler, nil
}