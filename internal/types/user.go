package types

type User struct {
	Name         string `json:"name"`
	Password     []byte `json:"password"`
	WebsocketKey string `json:"websocket_key"`
	Email        string `json:"email"`
}
