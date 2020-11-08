package datasource

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/volatiletech/sqlboiler/boil"
	"time"
)

type Domain struct {
	IDBase64 string    `json:"id"`
	IDUuid   uuid.UUID `json:"-"`

	Name string `json:"name"`
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

func DomainsAll(ctx *gin.Context) ([]Domain, int64, error){
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

func (model *Domain)ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(model.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} /*else if !nameRegexp.MatchString(domain.Name) {
		errorMap["name"] = []string{"must be a valid domain name"}
	}*/

	return errorMap
}

func (model *Domain)Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if model.dbModel == nil {
		insert = true
		model.dbModel = &datamodels_raw.Domain{}
		model.IDUuid, _ = uuid.NewV4()
		model.IDBase64 = UUIDToBase64(model.IDUuid)
		model.dbModel.ID = model.IDUuid.String()
		model.dbModel.Name = model.Name
	} else {
		modified := false

		if model.dbModel.Name != model.Name {
			modified = true
			model.dbModel.Name = model.Name
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

	model.fromDB(model.dbModel)
	return true, nil
}

func (model *Domain)Delete(ctx *gin.Context) (bool, error) {
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


func (model *Domain)fromDB(dbModel *datamodels_raw.Domain) {
	model.IDBase64 = UUIDStringToBase64(dbModel.ID)
	model.IDUuid = UUIDFromString(dbModel.ID)

	model.Name = dbModel.Name
	if dbModel.CreatedAt.Valid {
		model.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		model.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}

	model.dbModel = dbModel
}

func (model *Domain)toDB(dbModel *datamodels_raw.Domain) {
	if len(dbModel.ID) == 0 {
		uuid, _ := uuid.NewV4()
		dbModel.ID = uuid.String()
	}

	if len(model.Name) > 0 {
		dbModel.Name = model.Name
	}
}