package server

import (
	asciiart "ascii-art-web/source/ascii-art"
	"html/template"
	"net/http"
	"strings"
)

type ErrorPageMsg struct {
	ErrorCode string
	ErrorMsg  string
}

type ResultPage struct {
	Text   string
	Banner string
	Result string
}

// template to parse to all the functions
var allhandletemp = template.Must(template.ParseFiles("templates/index.html"))

// function for error  handler
func errorHandler(w http.ResponseWriter, r *http.Request, err *ErrorPageMsg) {
	errtemp := template.Must(template.ParseFiles("templates/error.html"))
	errtemp.Execute(w, err)
}

// function to handle the main function.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	//validating the path request
	if r.URL.Path != "/" {
		err := ErrorPageMsg{ErrorCode: "404", ErrorMsg: "PATH NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errorHandler(w, r, &err)
		return
	}
	// checking for request method
	if r.Method != http.MethodGet {
		err := ErrorPageMsg{ErrorCode: "405", ErrorMsg: "METHOD NOT ALLOWED"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorHandler(w, r, &err)
		return
	}
	allhandletemp.Execute(w, nil)
}

// function to handle the result paga
func ResultHandler(w http.ResponseWriter, r *http.Request) {
	// validating path request
	if r.URL.Path != "/ascii-art" {
		err := ErrorPageMsg{ErrorCode: "405", ErrorMsg: "PATH NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errorHandler(w, r, &err)
		return
	}
	//validating request method
	if r.Method != http.MethodPost {
		err := ErrorPageMsg{ErrorCode: "405", ErrorMsg: "METHOD NOT FOUND"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorHandler(w, r, &err)
	}
	// validating the parse form
	err := r.ParseForm()
	if err != nil {
		err := ErrorPageMsg{ErrorCode: "500", ErrorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errorHandler(w, r, &err)
		return
	}
	// checking the input
	input := r.PostFormValue("input-text")
	checkedinput := strings.ReplaceAll(input, "\r\n", "")
	for _, char := range checkedinput {
		if char < 32 || char > 126 {
			err := ErrorPageMsg{ErrorCode: "400", ErrorMsg: "INVALID INPUT"}
			w.WriteHeader(http.StatusBadRequest)
			errorHandler(w, r, &err)
			return
		}
	}
	//validating the banners
	banner := r.PostFormValue("banner")
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		err := ErrorPageMsg{ErrorCode: "404", ErrorMsg: "BANNER NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errorHandler(w, r, &err)
		return
	}
	//call the generated asciiart function
	ascii, err := asciiart.GenerateArt(input, banner)
	if err != nil {
		err := ErrorPageMsg{ErrorCode: "500", ErrorMsg: "SERVER INTERNAL ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errorHandler(w, r, &err)
		return
	}

	output := ResultPage{Text: input, Banner: banner, Result: ascii}
	err = allhandletemp.Execute(w, output)
}
