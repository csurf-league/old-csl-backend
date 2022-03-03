package websocket

import (
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
func newClient(steamid string, s *websocket.Conn, r *Room) *Client {
	return &Client{
		SteamID: steamid,
		socket:  s,
		room:    r,
		send:    make(chan []byte, messageBufferSize),
	}
}

// Read client messages (from frontend (either from current room or hub))
func (c *Client) read(fromHub bool) {
	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			break
		}
		if fromHub {
			Hub.forward <- msg
		} else {
			c.room.forward <- msg
		}
	}
	c.socket.Close()
}

// Send message to client (to frontend)
func (c *Client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

// Removes a client from the room
func (c *Client) DeleteFromRoom(r *Room) {
	r.Clients = removeClient(r.Clients, c.GetIndex(r.Clients))
	close(c.send)
	c.socket.Close()
	c.room = nil
}

// Removes a client from the hub
func (c *Client) DeleteFromHub() {
	Hub.Clients = removeClient(Hub.Clients, c.GetIndex(Hub.Clients))
	close(c.send)
	c.socket.Close()
}

// Removes a client from slice by its index
func removeClient(s []*Client, i int) []*Client {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
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
