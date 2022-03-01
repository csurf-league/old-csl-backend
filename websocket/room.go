package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Room represents a single chat room
type Room struct {
	id         int
	maxplayers int
	clients    map[*Client]bool
	forward    chan []byte
	join       chan *Client
	leave      chan *Client
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		// TODO: change origin to frontend url
		return true
	},
}

// Create a new chat room
func NewRoom(id int, maxplayers int) *Room {
	return &Room{
		id:         id,
		maxplayers: maxplayers,
		clients:    make(map[*Client]bool),
		forward:    make(chan []byte),
		join:       make(chan *Client),
		leave:      make(chan *Client),
	}
}

// Run chat room and wait for actions
func (r *Room) Run() {
	log.Printf("running chat room %d", r.id)
	for {
		select {
		case client := <-r.join:
			r.joinRoom(client)
		case client := <-r.leave:
			r.leaveRoom(client)
		case msg := <-r.forward:
			r.printToChatAll(msg)
		}
	}
}

// Client joins the room
func (r *Room) joinRoom(c *Client) {
	log.Printf("new client in room %v", r.id)
	r.clients[c] = true
}

// Client leaves the room
func (r *Room) leaveRoom(c *Client) {
	log.Printf("client leaving room %v", r.id)
	delete(r.clients, c)
	close(c.send)
}

// Print message to all in the current room
func (r *Room) printToChatAll(msg []byte) {
	data := FromJSON(msg)
	log.Printf("[room %v] %v: %v", r.id, data.Sender, data)

	for client := range r.clients {
		select {
		case client.send <- msg:
			// success
		default:
			// not sure if this is possible/reachable but yeah
			delete(r.clients, client)
			close(client.send)
		}
	}
}

// Returns true if room is full
func (room *Room) IsFull() bool {
	return len(room.clients) == room.maxplayers
}
