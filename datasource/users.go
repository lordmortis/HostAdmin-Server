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
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	OldPassword string `json:"current_password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	PasswordConfirmation string `json:"password_confirmation,omitempty"`
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

func UsersAll(ctx *gin.Context) ([]User, error){
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

	dbModels, err := datamodels_raw.Users().All(ctx, dbCon)
	if err != nil {
		return nil, err
	}

	viewModels := make([]User, len(dbModels))
	for index := range dbModels {
		viewModel := User{}
		viewModel.FromDB(dbModels[index])
		viewModels[index] = viewModel
	}

	return viewModels, nil
}

func (user *User)FromDB(dbModel *datamodels_raw.User) {
	user.uuid = UUIDFromString(dbModel.ID)
	user.dbModel = dbModel

	user.ID = UUIDToBase64(user.uuid)
	user.Username = dbModel.Username
	user.Email = dbModel.Email

	if dbModel.CreatedAt.Valid {
		user.CreatedAt = dbModel.CreatedAt.Time.Format(time.RFC3339)
	}

	if dbModel.UpdatedAt.Valid {
		user.UpdatedAt = dbModel.UpdatedAt.Time.Format(time.RFC3339)
	}
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
		_ = user.SetPassword(user.NewPassword)
	}
}

func (user *User)Update(ctx *gin.Context) (bool, error) {
	modified := false
	insert := false
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)
	if user.dbModel == nil {
		insert = true
		modified = true
		user.dbModel = &datamodels_raw.User{}
		user.uuid, _ = uuid.NewV4()
		user.ID = UUIDToBase64(user.uuid)
		user.dbModel.ID = user.uuid.String()
	}

	if user.dbModel.Email != user.Email {
		modified = true
		user.dbModel.Email = user.Email
	}

	if user.dbModel.Username != user.Username {
		modified = true
		user.dbModel.Username = user.Username
	}

	if len(user.NewPassword) > 0 && len(user.PasswordConfirmation) > 0 && user.NewPassword == user.PasswordConfirmation {
		modified = true
		user.SetPassword(user.NewPassword)
	}

	if !modified {
		return false, nil
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

	user.FromDB(user.dbModel)
	return modified, nil
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
	dbCon := ctx.MustGet("databaseConnection").(*sql.DB)

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

func UserWithIDString(ctx *gin.Context, stringID string) (*User, error) {
	userID := UUIDFromString(stringID)
	return UserWithUUID(ctx, userID)
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
	model.FromDB(dbModel)

	return &model, nil
}