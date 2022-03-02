package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

// Client represents a... client
type Client struct {
	SteamID string `json:"steamid"`
	//name    string `json:"name"`
	socket *websocket.Conn
	room   *Room
	send   chan []byte
}

// Create new client
func newClient(steamID string, s *websocket.Conn, r *Room) *Client {
	return &Client{
		SteamID: steamID,
		socket:  s,
		room:    r,
		send:    make(chan []byte, messageBufferSize),
	}
}

// Removes a client from slice by its index
func removeClient(s []*Client, i int) []*Client {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Read client messages (from frontend)
func (c *Client) read() {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		c.room.forward <- msg
	}
	c.socket.Close()
}

// Send message to client (to frontend)
func (c *Client) write() {
	for msg := range c.send {
		log.Println(string(msg))
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

// Returns client index on the slice
func (c *Client) GetIndex(slice []*Client) int {
	for idx, client := range slice {
		if client == c {
			return idx
		}
	}
	return -1
}
