package main

import (
	"ascii-art-web/source/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("server is running on a;8080\n Press ctl+c to stop the server.")
	//created a new servemux
	mux := http.NewServeMux()

	mux.HandleFunc("/", server.Hello)
	serverAdress := ":8010"

	err := http.ListenAndServe(serverAdress, mux)
	//check when server fails to start
	if err != nil {
		log.Fatalf("Server failds to start: %v", err)
	}
}
