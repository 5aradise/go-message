package database

import "github.com/5aradise/go-message/internal/types"

func UserDBToTypes(uDB User) types.User {
	return types.User{
		Name: uDB.Name,
		Password: uDB.Password,
		WebsocketKey: uDB.WebsocketKey,
		Email: uDB.Email.String,
	}
}