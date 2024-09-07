package database

import (
	"database/sql"

	"github.com/5aradise/go-message/internal/types"
	"github.com/5aradise/go-message/pkg/random"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string         `gorm:"unique"`
	Password     []byte         `gorm:"not null"`
	Email        sql.NullString `gorm:"unique"`
	WebsocketKey string         `gorm:"not null"`
}

func (sl *Database) CreateUser(name string, password []byte, email sql.NullString) (types.User, error) {
	wsKey, err := random.String(64)
	if err != nil {
		return types.User{}, err
	}

	user := User{
		Name:         name,
		Password:     password,
		Email:        email,
		WebsocketKey: wsKey,
	}
	res := sl.gormDB.Create(&user)
	return UserDBToTypes(user), res.Error
}

func (sl *Database) GetUserByName(name string) (types.User, error) {
	var user User
	res := sl.gormDB.Where(&User{Name: name}).First(&user)
	return UserDBToTypes(user), res.Error
}

func (sl *Database) GetUserByWebsocketKey(wsKey string) (types.User, error) {
	var user User
	res := sl.gormDB.Where(&User{WebsocketKey: wsKey}).First(&user)
	return UserDBToTypes(user), res.Error
}
