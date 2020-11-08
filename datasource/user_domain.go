package datasource

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"time"
)

type UserDomain struct {
	UserID string `json:"userID,omitempty""`
	User *User `json:"user,omitempty"`
	DomainID string `json:"domainID,omitempty"`
	Domain *Domain `json:"domain,omitempty"`

	DomainAdmin bool `json:"admin"`
	EmailAdmin bool `json:"email"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.UserDomain
}

func (domain *Domain)Users(ctx *gin.Context, loadUsers bool) ([]UserDomain, int64, error) {
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

func (model *UserDomain)fromDB(ctx *gin.Context, dbCon *sql.DB, dbModel *datamodels_raw.UserDomain, populateUser bool , populateDomain bool) {
	if populateUser {
		dbUser, _ := dbModel.User().One(ctx, dbCon)
		user := User{}
		user.fromDB(dbUser)
		model.User = &user
	} else {
		model.UserID = UUIDStringToBase64(dbModel.UserID)
	}

	if populateDomain {
		dbDomain, _ := dbModel.Domain().One(ctx, dbCon)
		domain := Domain{}
		domain.fromDB(dbDomain)
		model.Domain = &domain
	} else {
		model.DomainID = UUIDStringToBase64(dbModel.DomainID)
	}

	model.DomainAdmin = dbModel.Admin
	model.EmailAdmin = dbModel.Email

	if dbModel.CreatedAt.Valid {
		model.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		model.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}



	model.dbModel = dbModel
}
