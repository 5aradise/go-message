package types

type Message struct {
	Sender string `json:"sender"`
	Body   []byte `json:"body"`
}
