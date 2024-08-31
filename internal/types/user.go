package types

import (
	"iter"

	"github.com/gorilla/websocket"
)

type User struct {
	Name   string
	Key    string
	WsConn *websocket.Conn
}

type UserDB interface {
	SetUser(string, *User)
	GetUserByName(string) (*User, bool)
	GetUserByKey(string) (*User, bool)
	DeleteUser(string)
}

type UsersIterator interface {
	IterUsers() iter.Seq2[string, *User]
}
