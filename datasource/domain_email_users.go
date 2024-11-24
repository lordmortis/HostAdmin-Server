package datasource

import (
	"database/sql"
	"github.com/volatiletech/null/v8"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gopkg.in/errgo.v2/errors"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

type DomainEmailUser struct {
	DomainID             string  `json:"domainID,omitempty"`
	Domain               *Domain `json:"domain,omitempty"`
	BaseAddress          string  `json:"baseAddress,omitempty"`
	Password             string  `json:"password,omitempty"`
	PasswordConfirmation string  `json:"passwordConfirmation,omitempty"`
	Enabled              bool    `json:"enabled"`
	Quota                int     `json:"quota,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.DomainEmailUser
}

func (domain *Domain) EmailUsers(ctx *gin.Context) ([]DomainEmailUser, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	countQuery := domain.dbModel.DomainEmailUsers()
	dataQuery := domain.dbModel.DomainEmailUsers()

	count, err := countQuery.Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	if count == 0 {
		models := make([]DomainEmailUser, 0)
		return models, count, nil
	}

	dbModels, err := dataQuery.All(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	models := make([]DomainEmailUser, len(dbModels))
	for index := range dbModels {
		model := DomainEmailUser{}
		model.fromDB(ctx, dbCon, dbModels[index], false)
		models[index] = model
	}

	return models, count, nil
}

func DomainEmailUsers(ctx *gin.Context, domainID uuid.UUID, baseAddress string, populateDomain bool) (*DomainEmailUser, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, err
	}

	dbModel, err := datamodels_raw.FindDomainEmailUser(ctx, dbCon, domainID.String(), baseAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := DomainEmailUser{}
	model.fromDB(ctx, dbCon, dbModel, populateDomain)

	return &model, nil
}

func (model *DomainEmailUser) Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if model.dbModel == nil {
		insert = true
		dbModel := datamodels_raw.DomainEmailUser{}
		model.dbModel = &dbModel

		if model.Domain != nil {
			dbModel.DomainID = model.Domain.dbModel.ID
		} else {
			dbModel.DomainID = model.DomainID
		}

		dbModel.BaseAddress = model.BaseAddress

		bytes, err := argonHash(model.Password)
		if err != nil {
			return false, errors.Because(err, nil, "could not generate hashed password")
		}
		dbModel.EncryptedPassword = null.BytesFrom(bytes)

		dbModel.Enabled = model.Enabled
		dbModel.Quota = model.Quota
	} else {
		modified := false
		dbModel := model.dbModel
		if dbModel.BaseAddress != model.BaseAddress {
			modified = true
			dbModel.BaseAddress = model.BaseAddress
		}

		if len(model.Password) > 0 {
			modified = true
			bytes, err := argonHash(model.Password)
			if err != nil {
				return false, errors.Because(err, nil, "could not generate hashed password")
			}
			dbModel.EncryptedPassword = null.BytesFrom(bytes)
		}

		if dbModel.Enabled != model.Enabled {
			modified = true
			dbModel.Enabled = model.Enabled
		}

		if dbModel.Quota != model.Quota {
			modified = true
			dbModel.Quota = model.Quota
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

func (model *DomainEmailUser) Delete(ctx *gin.Context) (bool, error) {
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

func (model *DomainEmailUser) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if model.dbModel == nil && model.Domain == nil {
		if model.DomainID == "" {
			errorMap["domainID"] = []string{"must be present"}
		} else if UUIDFromString(model.DomainID) == uuid.Nil {
			errorMap["domainID"] = []string{"must be valid UUID"}
		}
	}

	if model.dbModel == nil && len(model.BaseAddress) == 0 {
		errorMap["baseAddress"] = []string{"must be present"}
	} else if model.dbModel != nil && model.BaseAddress != model.dbModel.BaseAddress {
		errorMap["baseAddress"] = []string{"cannot be modified"}
	}

	if model.dbModel == nil || len(model.Password) > 0 {
		passwordErrors := make([]string, 0, 2)
		passwordConfirmErrors := make([]string, 0, 2)

		if len(model.Password) == 0 {
			passwordErrors = append(passwordErrors, "must be present")
		}

		if len(model.PasswordConfirmation) == 0 {
			passwordConfirmErrors = append(passwordConfirmErrors, "must be present")
		}

		if model.PasswordConfirmation != model.Password {
			passwordErrors = append(passwordErrors, "must be equal")
			passwordConfirmErrors = append(passwordConfirmErrors, "must be equal")
		}

		if len(passwordErrors) > 0 {
			errorMap["password"] = passwordErrors
		}

		if len(passwordConfirmErrors) > 0 {
			errorMap["passwordConfirmation"] = passwordConfirmErrors
		}
	}

	if model.Quota == 0 {
		errorMap["quota"] = []string{"must be greater than zero"}
	}

	return errorMap
}

func (model *DomainEmailUser) fromDB(ctx *gin.Context, dbCon *sql.DB, dbModel *datamodels_raw.DomainEmailUser, populateDomain bool) {
	if populateDomain {
		dbDomain, _ := dbModel.Domain().One(ctx, dbCon)
		domain := Domain{}
		domain.fromDB(dbDomain)
		model.Domain = &domain
		model.DomainID = domain.IDBase64
	} else {
		model.DomainID = UUIDStringToBase64(dbModel.DomainID)
	}

	model.BaseAddress = dbModel.BaseAddress
	model.Enabled = dbModel.Enabled
	model.Quota = dbModel.Quota
	model.CreatedAt = dbModel.CreatedAt.Format(time.RFC3339)
	model.UpdatedAt = dbModel.UpdatedAt.Format(time.RFC3339)

	model.dbModel = dbModel
}
