package server

import (
	"fmt"
	"net/http"
)

// create a handle
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// w is for send data back and r is for reading incoming data
	fmt.Fprint(w, "Welcome to my Go server!")
}

func main() {
	//create a router
	mux := http.NewServeMux() 

	// the router is associated with the a handle
	mux.HandleFunc("/", homeHandler)

	// using router the engine is start on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
