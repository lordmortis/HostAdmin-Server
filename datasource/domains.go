package datasource

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
)

type Domain struct {
	IDBase64 string    `json:"id"`
	IDUuid   uuid.UUID `json:"-"`

	Name      string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.Domain
}

var (
//TODO: re-enable domain validation
//	nameRegexp *regexp.Regexp
)

func init() {
	//	nameRegexp = regexp.MustCompile("^(((?!-))(xn--|_{1,1})?[a-z0-9-]{0,61}[a-z0-9]{1,1}.)*(xn--)?([a-z0-9-]{1,61}|[a-z0-9-]{1,30}.[a-z]{2,})$")
}

func DomainsAll(ctx *gin.Context) ([]Domain, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	count, err := datamodels_raw.Domains().Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	dbModels, err := datamodels_raw.Domains().All(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	models := make([]Domain, len(dbModels))
	for index := range dbModels {
		model := Domain{}
		model.fromDB(dbModels[index])
		models[index] = model
	}

	return models, count, nil
}

func DomainWithID(ctx *gin.Context, id uuid.UUID) (*Domain, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, err
	}

	dbModel, err := datamodels_raw.FindDomain(ctx, dbCon, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := Domain{}
	model.fromDB(dbModel)

	return &model, nil
}

func (domain *Domain) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(domain.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} /*else if !nameRegexp.MatchString(domain.Name) {
		errorMap["name"] = []string{"must be a valid domain name"}
	}*/

	return errorMap
}

func (domain *Domain) Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if domain.dbModel == nil {
		insert = true
		domain.dbModel = &datamodels_raw.Domain{}
		domain.IDUuid, _ = uuid.NewV4()
		domain.IDBase64 = UUIDToBase64(domain.IDUuid)
		domain.dbModel.ID = domain.IDUuid.String()
		domain.dbModel.Name = domain.Name
	} else {
		modified := false

		if domain.dbModel.Name != domain.Name {
			modified = true
			domain.dbModel.Name = domain.Name
		}

		if !modified {
			return false, nil
		}
	}

	if insert {
		err := domain.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}
	} else {
		rows, err := domain.dbModel.Update(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}

		if rows == 0 {
			return false, nil
		}
	}

	if err := domain.dbModel.Reload(ctx, dbCon); err != nil {
		return false, err
	}

	domain.fromDB(domain.dbModel)
	return true, nil
}

func (domain *Domain) Delete(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	rows, err := domain.dbModel.Delete(ctx, dbCon)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	}

	return true, nil
}

func (domain *Domain) fromDB(dbModel *datamodels_raw.Domain) {
	domain.IDBase64 = UUIDStringToBase64(dbModel.ID)
	domain.IDUuid = UUIDFromString(dbModel.ID)

	domain.Name = dbModel.Name
	domain.CreatedAt = dbModel.CreatedAt.Format(time.RFC3339)
	domain.UpdatedAt = dbModel.UpdatedAt.Format(time.RFC3339)

	domain.dbModel = dbModel
}

func (domain *Domain) toDB(dbModel *datamodels_raw.Domain) {
	if len(dbModel.ID) == 0 {
		uuid, _ := uuid.NewV4()
		dbModel.ID = uuid.String()
	}

	if len(domain.Name) > 0 {
		dbModel.Name = domain.Name
	}
}
