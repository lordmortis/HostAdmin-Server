package datasource

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func dbFromContext(ctx *gin.Context) (*sql.DB, error) {
	val, ok := ctx.Get("databaseConnection")
	if !ok {
		return nil, ErrNoDb
	}

	dbCon, ok := val.(*sql.DB)
	if !ok {
		return nil, ErrNoDb
	}

	return dbCon, nil
}
