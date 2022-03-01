package websocket

import "net/http"

// /api/rooms/{room_id} - create a client websocket, associating it with a room
func (room *Room) HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	// create client socket
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// create new client associated to a room
	client := newClient(socket, room)

	// join the room
	room.join <- client

	// executed at end of this function (after goroutines stop)
	defer func() {
		room.leave <- client
	}()

	// run write and read in 2 separate goroutines
	go client.write()
	client.read()
}
