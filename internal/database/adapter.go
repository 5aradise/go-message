package database

import "github.com/5aradise/go-message/internal/types"

func UserDBToTypes(uDB User) types.User {
	return types.User{
		Name:         uDB.Name,
		Password:     uDB.Password,
		Email:        uDB.Email.String,
		RefreshToken: uDB.RefreshToken.Token,
	}
}
