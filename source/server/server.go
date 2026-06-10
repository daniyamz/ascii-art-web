package server

import (
	asciiart "ascii-art-web/source/ascii-art"
	"html/template"
	"net/http"
)

type ResultPage struct {
	Result string
}

var maintemp = template.Must(template.ParseFiles("templates/index.html"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := maintemp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func ResultHandler(w http.ResponseWriter, r *http.Request) {
	input := r.PostFormValue("input-text")
	banner := r.PostFormValue("banner")
	if input == "" {
		http.Error(w, "Bad Request: Please enter text input", http.StatusBadRequest)
		return
	}
	ascii, err := asciiart.GenerateArt(input, banner)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}

	output := ResultPage{Result: ascii}
	err = maintemp.Execute(w, output)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
