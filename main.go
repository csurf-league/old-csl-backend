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

	// auth
	router.HandleFunc("/login",  controller.Login).Methods("GET")
	router.HandleFunc("/logout", controller.Logout).Methods("GET")
	router.Handle("/auth", 		 middleware.BeforeAuth(controller.Auth)).Methods("PUT", "OPTIONS");

	// steam_user (steam data)
	router.Handle("/profile", middleware.IsAuthenticated(controller.GetProfile))
	router.HandleFunc("/api/steam/{steamid}", controller.GetSteamUser).Methods("GET")

	// player_stats (game statistics)
	router.HandleFunc("/api/player/{steamid}", controller.GetPlayer).Methods("GET")

	// hub/lobby
	websocket.Hub = websocket.NewHub()
	router.HandleFunc("/api/rooms", websocket.HandleHub) //TODO: apply isauth middleware after everything is done
	go websocket.RunHub()

	// rooms (TODO: no hardcode range)
	for _, uid := range []int{1, 2, 3, 4, 5} {
		newRoom := websocket.NewRoom(uid, 3)
		websocket.AddRoomToHub(newRoom)
		router.HandleFunc("/api/room/"+fmt.Sprint(uid), newRoom.HandleRoom) //TODO: apply isauth middleware after everything is done
		go newRoom.Run()
	}
}
