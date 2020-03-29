package viewmodels

import "github.com/lordmortis/HostAdmin-Server/datamodels_raw"

type DomainUser struct {
	UserID string `json:"user_id"`
	Username string `json:"username"`
	Email string `json:"email"`
}

func (domainUser *DomainUser)FromDB(dbModel *datamodels_raw.UserDomain) {
}