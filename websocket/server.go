package websocket

import "log"

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	rooms      map[*Room]bool
}

var WS *WsServer

func NewWebsocketServer() *WsServer {
	wsServer := &WsServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[*Room]bool),
	}

	return wsServer
}

// Run the websocket server
func (sv *WsServer) Run() {
	for {
		select {

		case client := <-sv.register:
			sv.registerClient(client)

		case client := <-sv.unregister:
			sv.unregisterClient(client)
		}
	}
}

func (sv *WsServer) registerClient(client *Client) {
	sv.publishClientJoined(client)
	//sv.listOnlineClients(client)
	sv.clients[client] = true
}

func (sv *WsServer) unregisterClient(client *Client) {
	if _, ok := sv.clients[client]; ok {
		delete(sv.clients, client)
		sv.publishClientLeft(client)
	}
}

func (sv *WsServer) publishClientJoined(client *Client) {
	message := &Message{
		Action: "UserJoinedAction",
		Sender: client,
	}

	log.Println("Someone connected to the websocket")
	log.Println(message)
}

func (sv *WsServer) publishClientLeft(client *Client) {
	message := &Message{
		Action: "UserLeftAction",
		Sender: client,
	}

	log.Println("Someone disconnected from the websocket")
	log.Println(message)
}
