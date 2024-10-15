package types

import "net/http"

type ChatCreator interface {
	CreateChat(name string) error
}

type ChatConnecter interface {
	ConnectToChat(chatName, userName string, w http.ResponseWriter, r *http.Request) error
}

type UserDeleter interface {
	DeleteFromChat(chatName, userName string) error
}
