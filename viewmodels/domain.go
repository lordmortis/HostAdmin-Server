package viewmodels

import (
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"time"
)

type Domain struct {
	ID string `json:"id"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

var (
//TODO: re-enable domain validation
	//	nameRegexp *regexp.Regexp
)

func init() {
//	nameRegexp = regexp.MustCompile("^(((?!-))(xn--|_{1,1})?[a-z0-9-]{0,61}[a-z0-9]{1,1}.)*(xn--)?([a-z0-9-]{1,61}|[a-z0-9-]{1,30}.[a-z]{2,})$")
}


func (domain *Domain)FromDB(dbModel *datamodels_raw.Domain) {
	domain.ID = datamodels.UUIDFromString(dbModel.ID).String()
	domain.Name = dbModel.Name
	if dbModel.CreatedAt.Valid {
		domain.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		domain.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (domain *Domain)ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(domain.Name) == 0 {
		errorMap["name"] = []string{"must be present"}
	} /*else if !nameRegexp.MatchString(domain.Name) {
		errorMap["name"] = []string{"must be a valid domain name"}
	}*/


	return errorMap
}


func (domain *Domain)ToDB(dbModel *datamodels_raw.Domain) {
	if len(dbModel.ID) == 0 {
		uuid, _ := uuid.NewV4()
		dbModel.ID = uuid.String()
	}

	if len(domain.Name) > 0 {
		dbModel.Name = domain.Name
	}
}