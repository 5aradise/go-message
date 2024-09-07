package types

import "database/sql"

type UserCreator interface {
	CreateUser(name string, password []byte, email sql.NullString) (User, error)
}

type UserGetterByName interface {
	GetUserByName(name string) (User, error)
}

type UserGetterByWsKey interface {
	GetUserByWebsocketKey(wsKey string) (User, error)
}
