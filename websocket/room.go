package websocket

import (
	"fmt"
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
			r.broadcastToAll(msg)
		}
	}
}

// Client joins the room
func (r *Room) joinRoom(c *Client) {
	msg := NewMessage("join-room", fmt.Sprintf("%s has joined the room", c.SteamID), c.SteamID, "now").ToJSON()
	r.broadcastToAll(msg)
	Hub.forward <- msg // Hub needs to know so it updates the current rooms
	r.Clients = append(r.Clients, c)
}

// Client leaves the room
func (r *Room) leaveRoom(c *Client) {
	msg := NewMessage("left-room", fmt.Sprintf("%s has left the room", c.SteamID), c.SteamID, "now").ToJSON()
	r.broadcastToAll(msg)
	Hub.forward <- msg
	c.DeleteFromRoom(r)
}

// Sends a message to all clients connected to the room socket
func (r *Room) broadcastToAll(msg []byte) {
	for _, client := range r.Clients {
		select {
		case client.send <- msg:
			// success
		default:
			// not sure if this is possible/reachable but yeah
			client.DeleteFromRoom(r)
		}
	}
}

// Returns true if room is full
func (room *Room) IsFull() bool {
	return len(room.Clients) == room.Maxplayers
}
