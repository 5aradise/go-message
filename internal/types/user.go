package types

type User struct {
	Name         string `json:"name"`
	Password     []byte `json:"password"`
	Email        string `json:"email"`
	RefreshToken string `json:"refresh_token"`
	ChatName     string `json:"chat_name"`
}
