package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "server is working!")
	http.ServeContent()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
