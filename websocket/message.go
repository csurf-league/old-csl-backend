package websocket

import "encoding/json"

// Message represents a chat message
type Message struct {
	Message string `json:"message"`
	Sender  string `json:"sender"`
	Created string `json:"created"`
}

// FromJSON creates a new Message struct from given JSON
func FromJSON(jsonData []byte) (message *Message) {
	json.Unmarshal(jsonData, &message)
	return message
}
