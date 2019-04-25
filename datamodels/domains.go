package datamodels

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

func DomainById(ctx *gin.Context, dbCon *sql.DB, stringID string) (*datamodels_raw.Domain, error) {
	id := UUIDFromString(stringID)

	if id == uuid.Nil {
		return nil, UUIDParseError
	}

	user, err := datamodels_raw.FindDomain(ctx,dbCon, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}