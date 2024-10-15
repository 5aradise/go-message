package types

type Message struct {
	Sender string `json:"sender"`
	Body   string `json:"body"`
}

func NewMessage(sender, body string) Message {
	return Message{
		Sender: sender,
		Body:   body,
	}
}
