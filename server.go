package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robyzzz/csl-backend/controller"
	"github.com/robyzzz/csl-backend/middleware"
	"github.com/robyzzz/csl-backend/utils/env"
)

var router *mux.Router

func main() {
	env.GetEnvVariables()

	setupRouter()

	fmt.Printf("router initialized and listening on %s\n", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}

func setupRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/", controller.Home)
	router.Handle("/login", middleware.IsAuthenticated(controller.Login))
}
