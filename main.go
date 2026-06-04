package main

import (
	"ascii-art-web/source/server"
	"log"
	"net/http"
)

func main() {
	//created a new servemux
	mux := http.NewServeMux()

	mux.HandleFunc("/", server.Hello)
	serverAdress := ":8080"

	err := http.ListenAndServe(serverAdress, mux)
	//check when server fails to start
	if err != nil {
		log.Fatalf("Server failds to start: %v", err)
	}
}
