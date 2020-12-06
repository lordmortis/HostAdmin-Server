package datasource

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/lordmortis/HostAdmin-Server/datamodels_raw"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type User struct {
	ID *string `json:"id,omitempty"`
	Username *string `json:"username"`
	Email *string `json:"email"`
	SuperAdmin *bool `json:"super_admin,omitempty"`
	OldPassword *string `json:"current_password,omitempty"`
	NewPassword *string `json:"new_password,omitempty"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`

	uuid uuid.UUID
	dbModel *datamodels_raw.User
}

var (
	emailRegexp *regexp.Regexp
)

func init() {
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
}

func UsersWithUsername(ctx *gin.Context, username *string) (datamodels_raw.UserSlice, error){
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

	user.ID = UUIDToBase64Ptr(user.uuid)
	user.Username = &dbModel.Username
	user.Email = &dbModel.Email
	user.SuperAdmin = &dbModel.SuperAdmin.Bool

	if dbModel.CreatedAt.Valid {
		user.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		user.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
}

func (user *User) toDB(dbModel *datamodels_raw.User) {
	if len(dbModel.ID) == 0 {
		id, _ := uuid.NewV4()
		dbModel.ID = id.String()
	}

	if user.Email != nil {
		dbModel.Email = *user.Email
	}

	if user.SuperAdmin != nil {
		dbModel.SuperAdmin.Bool = *user.SuperAdmin
	}

	if user.Username != nil {
		dbModel.Username = *user.Username
	}

	if user.NewPassword != nil {
		_ = user.SetPassword(*user.NewPassword)
	}
}

func (user *User)ParseJSON(ctx *gin.Context) error {
	user.ID = nil
	user.Username = nil
	user.Email = nil
	user.SuperAdmin = nil
	user.OldPassword = nil
	user.NewPassword = nil
	user.PasswordConfirmation = nil
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return err
	}

	return nil
}

func (user *User)Update(ctx *gin.Context, admin bool) (bool, error) {
	dbCon, err := dbFromContext(ctx)
	if err != nil {
		return false, err
	}

	insert := false
	columns := boil.Infer()

	if user.dbModel == nil {
		insert = true
		user.dbModel = &datamodels_raw.User{}
		user.uuid, _ = uuid.NewV4()
		user.ID = UUIDToBase64Ptr(user.uuid)
		user.dbModel.ID = user.uuid.String()
		user.dbModel.Username = *user.Username
		user.dbModel.Email = *user.Email
		user.dbModel.SuperAdmin.Valid = true
		if user.SuperAdmin == nil {
			user.dbModel.SuperAdmin.Bool = false
		} else {
			user.dbModel.SuperAdmin.Bool = *user.SuperAdmin
		}

	} else {

		columnList := make([]string, 0, 4)

		if user.Email != nil && user.dbModel.Email != *user.Email {
			columnList = append(columnList, datamodels_raw.UserColumns.Email)
			user.dbModel.Email = *user.Email
		}

		if user.Username != nil && user.dbModel.Username != *user.Username {
			columnList = append(columnList, datamodels_raw.UserColumns.Username)
			user.dbModel.Username = *user.Username
		}

		if user.NewPassword != nil {
			if !admin && !user.ValidatePassword(*user.OldPassword) {
				return false, ErrUnauthorized
			}
			columnList = append(columnList, datamodels_raw.UserColumns.EncryptedPassword)
			user.SetPassword(*user.NewPassword)
		}

		if admin && user.SuperAdmin != nil {
			columnList = append(columnList, datamodels_raw.UserColumns.SuperAdmin)
			user.dbModel.SuperAdmin.Valid = true
			user.dbModel.SuperAdmin.Bool = *user.SuperAdmin
		}

		if len(columnList) == 0 {
			return false, nil
		}

		columns = boil.Whitelist(columnList...)
	}

	if insert {
		err := user.dbModel.Insert(ctx, dbCon, columns)
		if err != nil {
			return false, err
		}
	} else {
		rows, err := user.dbModel.Update(ctx, dbCon, columns)
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

	if user.Username == nil {
		errorMap["username"] = []string{"must be present"}
	} else if len(*user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if user.Email == nil {
		errorMap["email"] = []string{"must be present"}
	} else if !emailRegexp.MatchString(*user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if user.NewPassword == nil {
		errorMap["new_password"] = []string{"required"}
		errorMap["password_confirmation"] = []string{"required"}
	} else if *user.NewPassword != *user.PasswordConfirmation {
		errorMap["new_password"] = []string{"must equal password_confirmation"}
		errorMap["password_confirmation"] = []string{"must equal new_password"}
	}

	return errorMap
}

func (user *User) ValidateUpdate(admin bool) map[string]interface{} {
	errorMap := make(map[string]interface{})

	if user.Username != nil && len(*user.Username) < 4 {
		errorMap["username"] = []string{"must be at least 4 characters"}
	}

	if user.Email != nil && !emailRegexp.MatchString(*user.Email) {
		errorMap["email"] = []string{"must be a valid email address"}
	}

	if user.NewPassword != nil {
		if !admin && user.OldPassword == nil {
			errorMap["old_password"] = []string{"must pass the old password to update password"}
		}

		if *user.NewPassword != *user.PasswordConfirmation {
			errorMap["new_password"] = []string{"must equal password_confirmation"}
			errorMap["password_confirmation"] = []string{"must equal new_password"}
		}
	}

	return errorMap
}

func (user *User)SetPassword(newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.dbModel.EncryptedPassword = null.BytesFrom(hashedPassword)
	return nil
}

func (user *User)ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword(user.dbModel.EncryptedPassword.Bytes, []byte(password)); err != nil {
		return false
	}
	return true
}

func (user *User)Delete(ctx *gin.Context) (bool, error) {
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

	dbModel, err := datamodels_raw.FindUser(ctx,dbCon, id.String())
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