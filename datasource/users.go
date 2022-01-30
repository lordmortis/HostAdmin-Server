package datasource

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   string `json:"id"`
	Username             string `json:"username"`
	Email                string `json:"email"`
	SuperAdmin           bool   `json:"superAdmin"`
	OldPassword          string `json:"currentPassword,omitempty"`
	NewPassword          string `json:"newPassword,omitempty"`
	PasswordConfirmation string `json:"passwordConfirmation,omitempty"`
	CreatedAt            string `json:"created_at,omitempty"`
	UpdatedAt            string `json:"updated_at,omitempty"`

	uuid    uuid.UUID
	dbModel *datamodels_raw.User
}

var (
	emailRegexp *regexp.Regexp
)

func init() {
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

func UsersWithUsername(ctx *gin.Context, username *string) (datamodels_raw.UserSlice, error) {
	dbCon, err := dbFromContext(ctx)

	if err != nil {
		return nil, err
	}

	models, err := datamodels_raw.Users(qm.Where("username = ?", username)).All(ctx, dbCon)
	if err == sql.ErrNoRows {
		return nil, ErrNoResults
	}

	return models, err
}

func UsersAll(ctx *gin.Context) ([]User, int64, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, -1, err
	}

	count, err := datamodels_raw.Users().Count(ctx, dbCon)
	if err != nil {
		return nil, -1, err
	}

	dbModels, err := datamodels_raw.Users().All(ctx, dbCon)
	if err != nil {
		return nil, count, err
	}

	viewModels := make([]User, len(dbModels))
	for index := range dbModels {
		viewModel := User{}
		viewModel.fromDB(dbModels[index])
		viewModels[index] = viewModel
	}

	return viewModels, count, nil
}

func (user *User) fromDB(dbModel *datamodels_raw.User) {
	user.uuid = UUIDFromString(dbModel.ID)
	user.dbModel = dbModel

	user.ID = UUIDToBase64(user.uuid)
	user.Username = dbModel.Username
	user.Email = dbModel.Email

	user.CreatedAt = dbModel.CreatedAt.Format(time.RFC3339)
	user.UpdatedAt = dbModel.UpdatedAt.Format(time.RFC3339)
}

func (user *User) toDB(dbModel *datamodels_raw.User) {
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
		_ = user.SetPassword(user.NewPassword)
	}
}

func (user *User) Update(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false

	if user.dbModel == nil {
		insert = true
		dbModel := datamodels_raw.User{}
		user.dbModel = &dbModel
		user.uuid, _ = uuid.NewV4()
		user.ID = UUIDToBase64(user.uuid)
		user.dbModel.ID = user.uuid.String()
		user.SetPassword(user.NewPassword)
		dbModel.Username = user.Username
		dbModel.Email = user.Email
		dbModel.SuperAdmin = user.SuperAdmin
	} else {
		modified := false

		if user.dbModel.Email != user.Email {
			modified = true
			user.dbModel.Email = user.Email
		}

		if user.dbModel.Username != user.Username {
			modified = true
			user.dbModel.Username = user.Username
		}

		if len(user.NewPassword) > 0 {
			modified = true
			user.SetPassword(user.NewPassword)
		}

		if user.dbModel.SuperAdmin != user.SuperAdmin {
			modified = true
			user.dbModel.SuperAdmin = user.SuperAdmin
		}

		if !modified {
			return false, nil
		}
	}

	if insert {
		err := user.dbModel.Insert(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}
	} else {
		rows, err := user.dbModel.Update(ctx, dbCon, boil.Infer())
		if err != nil {
			return false, err
		}

		if rows == 0 {
			return false, nil
		}
	}

	if err := user.dbModel.Reload(ctx, dbCon); err != nil {
		return false, err
	}

	user.fromDB(user.dbModel)
	return true, nil
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
		errorMap["newPassword"] = []string{"required"}
		errorMap["passwordConfirmation"] = []string{"required"}
	} else if user.NewPassword != user.PasswordConfirmation {
		errorMap["newPassword"] = []string{"must equal password_confirmation"}
		errorMap["passwordConfirmation"] = []string{"must equal new_password"}
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
		errorMap["newPassword"] = []string{"must equal password_confirmation"}
		errorMap["passwordConfirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) SetPassword(newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.dbModel.EncryptedPassword = hashedPassword
	return nil
}

func (user *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword(user.dbModel.EncryptedPassword, []byte(password)); err != nil {
		return false
	}
	return true
}

func (user *User) Delete(ctx *gin.Context) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	rows, err := user.dbModel.Delete(ctx, dbCon)
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func UserWithUUID(ctx *gin.Context, id uuid.UUID) (*User, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if id == uuid.Nil {
		return nil, UUIDParseError
	}

	dbModel, err := datamodels_raw.FindUser(ctx, dbCon, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	model := User{}
	model.fromDB(dbModel)

	return &model, nil
}
