package websocket

import (
	"encoding/json"
	"fmt"
	"log"
)

type RoomsHub struct {
	Rooms   []*Room   `json:"rooms"`
	Clients []*Client `json:"clients"`
	forward chan []byte
	join    chan *Client
	leave   chan *Client
}

var Hub *RoomsHub

// Create a new lobby room
func NewHub() *RoomsHub {
	return &RoomsHub{
		Rooms:   nil,
		Clients: nil,
		forward: make(chan []byte),
		join:    make(chan *Client),
		leave:   make(chan *Client),
	}
}

// Run hub and update by only sending actions to frontend
func RunHub() {
	for {
		select {
		case client := <-Hub.join:
			addClientToHub(client)

		case client := <-Hub.leave:
			removeClientFromHub(client)

		case msg := <-Hub.forward:
			handleHubMessage(msg)
		}
	}
}

// Client joins the hub
func addClientToHub(c *Client) {
	Hub.Clients = append(Hub.Clients, c)
	msg := NewMessage("join-hub", fmt.Sprintf("%s has joined the hub", c.SteamID), c.SteamID, "now").ToJSON()
	hubBroadcastToAll(msg)
	c.hubBroadcastToClient(GetRoomsJSON())
	c.hubBroadcastToClient(GetHubPlayersJSON())
}

// Client leaves the hub
func removeClientFromHub(c *Client) {
	c.DeleteFromHub()
	msg := NewMessage("left-hub", fmt.Sprintf("%s has left the hub", c.SteamID), c.SteamID, "now").ToJSON()
	hubBroadcastToAll(msg)
}

// Returns current rooms info
func GetRoomsJSON() []byte {
	data := struct {
		Action string  `json:"action"`
		Data   []*Room `json:"data"`
	}{
		Action: "get-rooms",
		Data:   Hub.Rooms,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	return response
}

// Returns current hub players (IN THE HUB, NOT ON ANY ROOM)
func GetHubPlayersJSON() []byte {
	data := struct {
		Action string    `json:"action"`
		Data   []*Client `json:"data"`
	}{
		Action: "get-hub-players",
		Data:   Hub.Clients,
	}
	response, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
	}
	return response
}

// Depending on the action, send a response back to client (frontend)
func handleHubMessage(msg []byte) {
	data := FromJSON(msg)
	var response []byte

	// TODO: find a better way for this type of communication (especially join and left room updates)
	switch data.Action {
	// update rooms:
	case "join-room", "left-room":
		response = GetRoomsJSON()
	}

	hubBroadcastToAll(response)
}

// Send message to all from current hub
func hubBroadcastToAll(msg []byte) {
	for _, client := range Hub.Clients {
		client.hubBroadcastToClient(msg)
	}
}

// Sends a hub message to a single client
func (c *Client) hubBroadcastToClient(msg []byte) {
	select {
	case c.send <- msg:
		// success
	default:
		// not sure if this is possible/reachable but yeah
		c.DeleteFromHub()
	}
}

// Add room to hub
func AddRoomToHub(r *Room) {
	Hub.Rooms = append(Hub.Rooms, r)
}

// Returns true if player is already in another room
func AlreadyInAnotherRoom(steamid string) bool {
	for _, r := range Hub.Rooms {
		for _, c := range r.Clients {
			if c.SteamID == steamid {
				return true
			}
		}
	}
	return false
}

// Returns true if player already is in the current hub
func AlreadyInHub(steamid string) bool {
	for _, c := range Hub.Clients {
		if c.SteamID == steamid {
			return true
		}
	}
	return false
}
