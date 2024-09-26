package database

import (
	"database/sql"
	"time"

	"github.com/5aradise/go-message/internal/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string         `gorm:"unique"`
	Password     []byte         `gorm:"not null"`
	Email        sql.NullString `gorm:"unique"`
	RefreshToken RefreshToken
}

func (sl *Database) CreateUser(name string, password []byte, email sql.NullString, refreshToken string) (types.User, error) {
	user := User{
		Name:     name,
		Password: password,
		Email:    email,
		RefreshToken: RefreshToken{
			Token:     refreshToken,
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		},
	}
	res := sl.gormDB.Create(&user)
	return UserDBToTypes(user), res.Error
}

func (sl *Database) GetUserByName(name string) (types.User, error) {
	var user User
	res := sl.gormDB.Where(&User{Name: name}).First(&user)
	return UserDBToTypes(user), res.Error
}

func (sl *Database) GetUserByRefreshToken(refreshToken string) (types.User, error) {
	var user User
	res := sl.gormDB.
		Where(&User{RefreshToken: RefreshToken{Token: refreshToken}}).
		First(&user)
	return UserDBToTypes(user), res.Error
}
