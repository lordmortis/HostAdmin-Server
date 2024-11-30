package datasource

import (
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gopkg.in/errgo.v2/errors"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

type DomainEmailAlias struct {
	DomainID     string   `json:"domainID,omitempty"`
	Domain       *Domain  `json:"domain,omitempty"`
	Address      string   `json:"address,omitempty"`
	Destinations []string `json:"destinations,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.DomainEmailAlias
}

func (domain *Domain) EmailAliases(ctx *gin.Context) ([]DomainEmailAlias, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	countQuery := domain.dbModel.DomainEmailAliases()
	dataQuery := domain.dbModel.DomainEmailAliases()

	count, err := countQuery.Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	if count == 0 {
		models := make([]DomainEmailAlias, 0)
		return models, count, nil
	}

	dbModels, err := dataQuery.All(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	models := make([]DomainEmailAlias, len(dbModels))
	for index := range dbModels {
		model := DomainEmailAlias{}
		model.fromDB(ctx, dbCon, dbModels[index], false)
		models[index] = model
	}

	return models, count, nil
}

func DomainEmailAliases(ctx *gin.Context, domainID uuid.UUID, address string, populateDomain bool) (*DomainEmailAlias, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, err
	}

	dbModel, err := datamodels_raw.FindDomainEmailAlias(ctx, dbCon, domainID.String(), address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := DomainEmailAlias{}
	model.fromDB(ctx, dbCon, dbModel, populateDomain)

	return &model, nil
}

func (model *DomainEmailAlias) Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if model.dbModel == nil {
		insert = true
		dbModel := datamodels_raw.DomainEmailAlias{}
		model.dbModel = &dbModel

		if model.Domain != nil {
			dbModel.DomainID = model.Domain.dbModel.ID
		} else {
			dbModel.DomainID = model.DomainID
		}

		if model.Destinations != nil {
			dbModel.Destinations = model.Destinations
		}

		dbModel.Address = model.Address
	} else {
		modified := false
		dbModel := model.dbModel
		if dbModel.Address != model.Address {
			modified = true
			dbModel.Address = model.Address
		}

		if model.Destinations != nil {
			modified = true
			dbModel.Destinations = model.Destinations
		}

		if !modified {
			return false, nil
		}
	}

	var rows int64

	if insert {
		rows = 1
		err = model.dbModel.Insert(ctx, dbCon, boil.Infer())
	} else {
		rows, err = model.dbModel.Update(ctx, dbCon, boil.Infer())
	}

	if err != nil {
		return false, errors.Because(err, nil, "unable to insert/update domain email user")
	}

	return rows > 0, nil
}

func (model *DomainEmailAlias) Delete(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	rows, err := model.dbModel.Delete(ctx, dbCon)
	if err != nil {
		return false, err
	}
	if rows == 0 {
		return false, nil
	}

	return true, nil
}

func (model *DomainEmailAlias) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if model.dbModel == nil && model.Domain == nil {
		if model.DomainID == "" {
			errorMap["domainID"] = []string{"must be present"}
		} else if UUIDFromString(model.DomainID) == uuid.Nil {
			errorMap["domainID"] = []string{"must be valid UUID"}
		}
	}

	if model.dbModel == nil && len(model.Address) == 0 {
		errorMap["address"] = []string{"must be present"}
	} else if model.dbModel != nil && model.Address != model.dbModel.Address {
		errorMap["address"] = []string{"cannot be modified"}
	}

	if model.dbModel == nil && (model.Destinations == nil || len(model.Destinations) == 0) {
		errorMap["destinations"] = []string{"must be present and have at least one value"}
	}

	return errorMap
}

func (model *DomainEmailAlias) fromDB(ctx *gin.Context, dbCon *sql.DB, dbModel *datamodels_raw.DomainEmailAlias, populateDomain bool) {
	if populateDomain {
		dbDomain, _ := dbModel.Domain().One(ctx, dbCon)
		domain := Domain{}
		domain.fromDB(dbDomain)
		model.Domain = &domain
		model.DomainID = domain.IDBase64
	} else {
		model.DomainID = UUIDStringToBase64(dbModel.DomainID)
	}

	model.Address = dbModel.Address
	model.Destinations = dbModel.Destinations
	model.CreatedAt = dbModel.CreatedAt.Format(time.RFC3339)
	model.UpdatedAt = dbModel.UpdatedAt.Format(time.RFC3339)

	model.dbModel = dbModel
}
