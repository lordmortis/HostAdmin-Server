package datasource

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

type UserDomain struct {
	UserID   string  `json:"userID,omitempty""`
	User     *User   `json:"user,omitempty"`
	DomainID string  `json:"domainID,omitempty"`
	Domain   *Domain `json:"domain,omitempty"`

	DomainAdmin bool `json:"domainAdmin"`
	EmailAdmin  bool `json:"emailAdmin"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.UserDomain
}

//goland:noinspection ALL
func (domain *Domain) Users(ctx *gin.Context, loadUsers bool) ([]UserDomain, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	countQuery := domain.dbModel.UserDomains()
	dataQuery := domain.dbModel.UserDomains()
	if loadUsers {
		dataQuery = domain.dbModel.UserDomains(qm.Load("User"))
	}

	count, err := countQuery.Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	if count == 0 {
		models := make([]UserDomain, 0)
		return models, count, nil
	}

	dbModels, err := dataQuery.All(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	models := make([]UserDomain, len(dbModels))
	for index := range dbModels {
		model := UserDomain{}
		model.fromDB(ctx, dbCon, dbModels[index], loadUsers, false)
		models[index] = model
	}

	return models, count, nil
}

//goland:noinspection ALL
func (parentModel *User) Domains(ctx *gin.Context, loadDomains bool) ([]UserDomain, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	countQuery := parentModel.dbModel.UserDomains()
	dataQuery := parentModel.dbModel.UserDomains()
	if loadDomains {
		dataQuery = parentModel.dbModel.UserDomains(qm.Load("Domain"))
	}

	count, err := countQuery.Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	if count == 0 {
		models := make([]UserDomain, 0)
		return models, count, nil
	}

	dbModels, err := dataQuery.All(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	models := make([]UserDomain, len(dbModels))
	for index := range dbModels {
		model := UserDomain{}
		model.fromDB(ctx, dbCon, dbModels[index], false, loadDomains)
		models[index] = model
	}

	return models, count, nil
}

func UserDomainsWithIDs(ctx *gin.Context, userID uuid.UUID, domainID uuid.UUID, populateUser bool, populateDomain bool) (*UserDomain, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, err
	}

	dbModel, err := datamodels_raw.FindUserDomain(ctx, dbCon, userID.String(), domainID.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := UserDomain{}
	model.fromDB(ctx, dbCon, dbModel, populateUser, populateDomain)

	return &model, nil
}

func (model *UserDomain) Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if model.dbModel == nil {
		insert = true
		model.dbModel = &datamodels_raw.UserDomain{}

		if model.Domain != nil {
			model.dbModel.DomainID = model.Domain.dbModel.ID
		} else {
			model.dbModel.DomainID = model.DomainID
		}

		if model.User != nil {
			model.dbModel.UserID = model.User.dbModel.ID
		} else {
			model.dbModel.UserID = model.UserID
		}

		model.dbModel.Admin = model.DomainAdmin
		model.dbModel.Email = model.EmailAdmin
	} else {
		modified := false
		if model.dbModel.Admin != model.DomainAdmin {
			modified = true
			model.dbModel.Admin = model.DomainAdmin
		}

		if model.dbModel.Email != model.EmailAdmin {
			modified = true
			model.dbModel.Email = model.EmailAdmin
		}

		if !modified {
			return false, nil
		}
	}

	if insert {
		err := model.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}
	} else {
		rows, err := model.dbModel.Update(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}

		if rows == 0 {
			return false, nil
		}
	}

	if err := model.dbModel.Reload(ctx, dbCon); err != nil {
		return false, err
	}

	model.fromDB(ctx, dbCon, model.dbModel, true, true)
	return true, nil
}

func (model *UserDomain) Delete(ctx *gin.Context) (bool, error) {
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

func (model *UserDomain) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if model.dbModel == nil && model.Domain == nil {
		if model.DomainID == "" {
			errorMap["domainID"] = []string{"must be present"}
		} else if UUIDFromString(model.DomainID) == uuid.Nil {
			errorMap["domainID"] = []string{"must be valid UUID"}
		}
	}

	if model.dbModel == nil && model.User == nil {
		if model.UserID == "" {
			errorMap["userID"] = []string{"must be present"}
		} else if UUIDFromString(model.UserID) == uuid.Nil {
			errorMap["userID"] = []string{"must be valid UUID"}
		}
	}

	return errorMap
}

func (model *UserDomain) fromDB(ctx *gin.Context, dbCon *sql.DB, dbModel *datamodels_raw.UserDomain, populateUser bool, populateDomain bool) {
	if populateUser {
		dbUser, _ := dbModel.User().One(ctx, dbCon)
		user := User{}
		user.fromDB(dbUser)
		model.User = &user
		model.UserID = user.ID
	} else {
		model.UserID = UUIDStringToBase64(dbModel.UserID)
	}

	if populateDomain {
		dbDomain, _ := dbModel.Domain().One(ctx, dbCon)
		domain := Domain{}
		domain.fromDB(dbDomain)
		model.Domain = &domain
		model.DomainID = domain.IDBase64
	} else {
		model.DomainID = UUIDStringToBase64(dbModel.DomainID)
	}

	model.DomainAdmin = dbModel.Admin
	model.EmailAdmin = dbModel.Email
	model.CreatedAt = dbModel.CreatedAt.Format(time.RFC3339)
	model.UpdatedAt = dbModel.UpdatedAt.Format(time.RFC3339)

	model.dbModel = dbModel
}
