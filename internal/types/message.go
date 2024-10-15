package types

import "time"

type Message struct {
	Sender    string    `json:"sender"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"time"`
}

func NewMessage(sender, body string, createdAt time.Time) Message {
	return Message{
		Sender:    sender,
		Body:      body,
		CreatedAt: createdAt.UTC(),
	}
}
