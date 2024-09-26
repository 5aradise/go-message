package types

import "database/sql"

type UserCreator interface {
	CreateUser(name string, password []byte, email sql.NullString, refreshToken string) (User, error)
}

type UserGetterByName interface {
	GetUserByName(name string) (User, error)
}

type UserGetterByRefreshToken interface {
	GetUserByRefreshToken(refreshToken string) (User, error)
}

type RefreshTokenUpdater interface {
	UpdateRefreshTokenByUserName(name string, newToken string) error
}
