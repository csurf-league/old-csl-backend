package websocket

import (
	"log"
)

// Room represents a single chat room
type Room struct {
	ID         int       `json:"id"`
	Maxplayers int       `json:"maxplayers"`
	Clients    []*Client `json:"clients"`
	forward    chan []byte
	join       chan *Client
	leave      chan *Client
}

// Create a new chat room
func NewRoom(id int, maxplayers int) *Room {
	return &Room{
		ID:         id,
		Maxplayers: maxplayers,
		Clients:    nil,
		forward:    make(chan []byte),
		join:       make(chan *Client),
		leave:      make(chan *Client),
	}
}

// Add room to hub
func AddRoomToHub(r *Room) {
	Lobby.Rooms = append(Lobby.Rooms, r)
}

// Run chat room and wait for actions
func (r *Room) Run() {
	log.Printf("running chat room %d", r.ID)
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
	log.Printf("new client (%v) in room %v", c.SteamID, r.ID)
	r.Clients = append(r.Clients, c)
}

// Client leaves the room
func (r *Room) leaveRoom(c *Client) {
	log.Printf("client (%v) leaving room %v", c.SteamID, r.ID)
	r.Clients = removeClient(r.Clients, c.GetIndex(r.Clients))
	close(c.send)

}

// Print message to all in the current room
func (r *Room) printToChatAll(msg []byte) {
	data := FromJSON(msg)
	log.Printf("[room %v] %v: %v", r.ID, data.Sender, data)

	for _, client := range r.Clients {
		select {
		case client.send <- msg:
			// success
		default:
			// not sure if this is possible/reachable but yeah
			r.Clients = removeClient(r.Clients, client.GetIndex(r.Clients))
			close(client.send)
		}
	}
}

// Returns true if room is full
func (room *Room) IsFull() bool {
	return len(room.Clients) == room.Maxplayers
}
