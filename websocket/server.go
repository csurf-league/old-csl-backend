package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		log.Println(string(msg))
		if err := conn.WriteMessage(msgType, msg); err != nil {
			log.Fatal(err.Error())
			return
		}
	}
}

// /api/rooms/{room_id} - create a client websocket, putting him in this room
func (room *Room) HandleRoom(w http.ResponseWriter, r *http.Request) {
	if room.IsFull() {
		// TODO: return error
		return
	}

	// create client socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		// TODO: return error
		return
	}

	// create new client associated to a room
	client := newClient("1337", socket, room)

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
