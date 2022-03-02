package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/robyzzz/csl-backend/utils"
)

type RoomsHub struct {
	Rooms []*Room `json:"rooms"`
}

var Lobby *RoomsHub

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

// /api/rooms
func RoomsWebsocket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create client socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO: return error
		return
	}

	log.Println("Someone connected to /rooms")

	reader(socket)
}

func reader(conn *websocket.Conn) {
	defer conn.Close()
	for {
		// read from frontend
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}

		data := FromJSON(msg)
		var response []byte
		if data.Action == "get-rooms" {
			response, _ = json.Marshal(Lobby.Rooms)
		}

		// send back msg
		if err := conn.WriteMessage(msgType, response); err != nil {
			log.Println(err.Error())
			return
		}
	}
}

// /api/rooms/{room_id} - create a client websocket, putting him in this room
func (room *Room) HandleRoom(w http.ResponseWriter, r *http.Request) {
	// TODO: check if client is not on another room

	if room.IsFull() {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusBadRequest, "Room is full"))
		return
	}

	// get steamid by url param TODO: change this to bearer auth?
	steamid := r.URL.Query().Get("steamid")
	if len(steamid) == 0 {
		log.Println("no steamid")
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusNotFound, "Invalid steamid"))
		return
	}

	// create client socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.APIErrorRespond(w, utils.NewAPIError(http.StatusInternalServerError, "Something went wrong..."))
		return
	}

	// create new client associated to a room
	client := newClient(steamid, socket, room)

	// join the room
	room.join <- client

	// executed at end of this fn
	defer func() {
		room.leave <- client
	}()

	// run write and read in 2 separate goroutines
	go client.write()
	client.read()
}
