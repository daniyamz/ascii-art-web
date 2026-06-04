package server

import (
	"net/http"
)

// function to hand index.html in the server.
func Hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
