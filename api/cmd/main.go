package main

import (
	"github.com/gorilla/mux"
)

func main() {
	handleRequest()
}

func handleRequest() {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("api/users").Methods("GET")
}
