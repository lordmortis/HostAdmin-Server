package datamodels

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID int64 `json:"id"`
	Username string
	HashedPassword []byte `json:"-"`
	CreatedAt	time.Time
}

func GeneratePassword(userPassword string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashed []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}