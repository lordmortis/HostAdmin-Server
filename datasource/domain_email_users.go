package datasource

import "github.com/lordmortis/HostAdmin-Server/datamodels_raw"

type DomainEmail struct {
	DomainID    string  `json:"domainID,omitempty"`
	Domain      *Domain `json:"domain,omitempty"`
	BaseAddress string  `json:"baseAddress,omitempty"`
	Password    string  `json:"password,omitempty"`
	Enabled     bool    `json:"enabled,omitempty"`
	Quota       int     `json:"quota,omitempty"`

	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	dbModel *datamodels_raw.DomainEmailUser
}
