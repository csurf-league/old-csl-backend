package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/utils"
)

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

// /api/rooms - handle the room lobby (aKa hub)
func HandleHub(w http.ResponseWriter, r *http.Request) {
	steamid := r.URL.Query().Get("steamid")
	if len(steamid) == 0 {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, "Invalid steamid"))
		return
	}

	if AlreadyInHub(steamid) {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusBadRequest, "Player already in hub"))
		return
	}

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusInternalServerError, "Something went wrong..."))
		return
	}

	for _, c := range r.Cookies() {
		fmt.Println(c)
	}

	log.Println("Someone connected to /rooms")

	client := newClient(steamid, socket, nil)
	Hub.join <- client

	defer func() {
		Hub.leave <- client
	}()

	go client.write()
	client.read(true)
}

// /api/room/{room_id} - create a client websocket, putting him in this room
func (room *Room) HandleRoom(w http.ResponseWriter, r *http.Request) {
	if room.IsFull() {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusBadRequest, "Room is full."))
		return
	}

	steamid := config.GetSessionID(r)
	if steamid == "" {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, "Invalid session ID."))
		return
	}

	if AlreadyInAnotherRoom(steamid) {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusBadRequest, "You can only join 1 room."))
		return
	}

	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusInternalServerError, "Something went wrong..."))
		return
	}

	client := newClient(steamid, socket, room)
	room.join <- client

	defer func() {
		room.leave <- client
	}()

	go client.write()
	client.read(false)
}
