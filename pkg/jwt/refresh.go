package jwt

import "github.com/5aradise/go-message/pkg/random"

func CreateRefreshToken() (string, error) {
	return random.String(64)
}
