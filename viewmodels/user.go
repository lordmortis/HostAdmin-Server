package viewmodels

import (
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/satori/go.uuid"
	"time"
)

type User struct {
	ID string `json:"id"`
	Username string
	Email string
	OldPassword string `json:"current_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

func (user *User)FromDB(dbModel *datamodels_raw.User) {
	user.ID = datamodels.UserUUIDBase64(dbModel)
	user.Username = dbModel.Username
	user.Email = dbModel.Email

	if dbModel.CreatedAt.Valid {
		user.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		user.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (user *User) Validate() bool {
	valid := true

	if len(user.Username) == 0 {
		valid = false
	}

	if len(user.Email) == 0 {
		valid = false
	}

	if len(user.NewPassword) > 0 && user.NewPassword != user.PasswordConfirmation {
		valid = false
	}

	return valid
}

func (user *User) ToDB(dbModel *datamodels_raw.User) {
	if len(dbModel.ID) == 0 {
		dbModel.ID = uuid.NewV4().String()
	}

	if len(user.Email) > 0 {
		dbModel.Email = user.Email
	}

	if len(user.Username) > 0 {
		dbModel.Username = user.Username
	}

	if len(user.NewPassword) > 0 && len(user.PasswordConfirmation) > 0 && user.NewPassword == user.PasswordConfirmation {
		_ = datamodels.UserSetPassword(dbModel, user.NewPassword)
	}
}
