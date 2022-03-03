package websocket

import (
	"encoding/json"
	"log"
)

// Message represents a chat message
type Message struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	Sender  string `json:"sender"`
	Created string `json:"created"`
}

// Message "constructor"
func NewMessage(a string, m string, s string, c string) *Message {
	return &Message{
		Action:  a,
		Message: m,
		Sender:  s,
		Created: c,
	}
}

// Converts a message to JSON
func (msg *Message) ToJSON() []byte {
	j, err := json.Marshal(msg)
	if err != nil {
		log.Println(err.Error())
	}
	return j
}

// FromJSON creates a new Message struct from given JSON
func FromJSON(jsonData []byte) (message *Message) {
	err := json.Unmarshal(jsonData, &message)
	if err != nil {
		log.Println(err.Error())
	}
	return message
}
