package viewmodels

import (
	"github.com/lordmortis/HostAdmin-Server/datamodels"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/gofrs/uuid"
	"regexp"
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

var (
	emailRegexp *regexp.Regexp
)

func init() {
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
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

func (user *User) ValidateCreate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(user.Username) == 0 {
		errorMap["username"] = []string{"must be present"}
	} else if len(user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if len(user.Email) == 0 {
		errorMap["email"] = []string{"must be present"}
	} else if !emailRegexp.MatchString(user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if len(user.NewPassword) == 0 {
		errorMap["new_password"] = []string{"required"}
		errorMap["password_confirmation"] = []string{"required"}
	} else if user.NewPassword != user.PasswordConfirmation {
		errorMap["new_password"] = []string{"must equal password_confirmation"}
		errorMap["password_confirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) ValidateUpdate() map[string]interface{} {
	errorMap := make(map[string]interface{})

	if len(user.Username) != 0 && len(user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if len(user.Email) > 0 && !emailRegexp.MatchString(user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if len(user.NewPassword) > 0 && user.NewPassword != user.PasswordConfirmation {
		errorMap["new_password"] = []string{"must equal password_confirmation"}
		errorMap["password_confirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) ToDB(dbModel *datamodels_raw.User) {
	if len(dbModel.ID) == 0 {
		id, _ := uuid.NewV4()
		dbModel.ID = id.String()
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
