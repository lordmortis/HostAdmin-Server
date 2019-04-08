package viewmodels

import (
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"time"
)

type User struct {
	ID string `json:"id"`
	Username string
	Email string
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func UserViewModel(dbModel *datamodels_raw.User) User {
	user := User{
		ID: datamodels.UserUUIDBase64(dbModel),
		Username: dbModel.Username,
		Email: dbModel.Email,
	}

	if dbModel.CreatedAt.Valid {
		user.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		user.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}

	return user
}