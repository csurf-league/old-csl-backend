package websocket

import "github.com/gorilla/websocket"

type Client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *Room
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
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
