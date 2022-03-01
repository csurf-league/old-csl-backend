package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

// Room represents a single chat room
type Room struct {
	id      uint
	forward chan []byte
	join    chan *Client
	leave   chan *Client
	clients map[*Client]bool
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: socketBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (room *Room) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create socket to client
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// create new user
	client := &Client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   room,
	}

	// join the room
	room.join <- client

	// executed at end of this function
	defer func() {
		room.leave <- client
	}()

	// run write and read function in 2 separate goroutines
	go client.write()
	client.read()
}

// Create a new chat room
func NewRoom(uid uint) *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		id:      uid,
	}
}

// Run chat room and wait for actions
func (r *Room) Run() {
	log.Printf("running chat room %d", r.id)
	for {
		select {
		case Client := <-r.join:
			r.joinRoom(Client)
		case Client := <-r.leave:
			r.leaveRoom(Client)
		case msg := <-r.forward:
			data := FromJSON(msg)
			log.Printf("Client '%v' writing message to room %v, message: %v", data.Sender, r.id, data.Message)

			// broadcast message to all
			for client := range r.clients {
				select {
				case client.send <- msg:
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
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
