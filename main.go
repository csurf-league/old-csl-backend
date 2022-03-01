package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robyzzz/csl-backend/config"
	"github.com/robyzzz/csl-backend/controller"
	"github.com/robyzzz/csl-backend/middleware"
	"github.com/robyzzz/csl-backend/model"
	"github.com/robyzzz/csl-backend/websocket"
)

var router *mux.Router

// start server
func main() {
	config.GetEnvVariables()
	model.Connect()

	setupRouter()

	log.Printf("Server initialized and listening on %s\n", config.PORT)
	log.Fatal(http.ListenAndServe(":"+config.PORT, router))
}

func setupRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/", controller.Home)
	router.Handle("/login", middleware.IsLogged(controller.Login))
	router.HandleFunc("/logout", controller.Logout)

	// steam user
	router.HandleFunc("/api/player/{steamid}", controller.GetSteamUser).Methods("GET")

	// player stats
	router.HandleFunc("/api/playerstats/{steamid}", controller.GetPlayerStats).Methods("GET")

	// create rooms and run them (TODO: database rooms integration)
	for _, uid := range []int{1, 2, 3, 4, 5} {
		newRoom := websocket.NewRoom(uid, 2)
		roomRoute := fmt.Sprintf("/chat/%d", uid)
		router.HandleFunc(roomRoute, newRoom.HandleWebsocket)
		go newRoom.Run()
	}
}
