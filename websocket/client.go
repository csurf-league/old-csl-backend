package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Client struct {
	conn     *websocket.Conn
	wsServer *WsServer
	send     chan []byte
	SteamID  string `json:"steamid"`
	Name     string `json:"name"`
	rooms    map[*Room]bool
}

func NewClient(conn *websocket.Conn, wsServer *WsServer, name string, steamID string) *Client {
	client := &Client{
		SteamID:  steamID,
		Name:     name,
		conn:     conn,
		wsServer: wsServer,
		send:     make(chan []byte, 256),
		rooms:    make(map[*Room]bool),
	}
	return client
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	// get params
	name := r.URL.Query().Get("name")
	steamid := r.URL.Query().Get("steamid")

	if len(name) == 0 || len(steamid) == 0 {
		log.Println("no name or no steamid")
		return
	}

	// receive websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("smt went wrong")
		return
	}

	client := NewClient(conn, WS, name, steamid)

	// handle I/O for client
	go client.readPump()
	go client.writePump()

	WS.register <- client
}

func (client *Client) readPump() {
	defer client.disconnect()

	//client.conn.SetReadLimit(maxMessageSize)
	//client.conn.SetReadDeadline(time.Now().Add(pongWait))
	//client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}

		client.handleNewMessage(jsonMessage)
	}
}

func (client *Client) writePump() {
	defer client.conn.Close()
	for {
		select {
		case message, ok := <-client.send:
			log.Println("write pump")
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write([]byte("\n"))
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		}
	}
}

func (client *Client) disconnect() {
	client.wsServer.unregister <- client
	for room := range client.rooms {
		room.unregister <- client
	}
	close(client.send)
	client.conn.Close()
}

func (client *Client) handleNewMessage(jsonMessage []byte) {
	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}

	log.Println("handle new message: ")
	log.Println(message)

	// send msg
	// join room
	// leave room
}
