package server

import (
	"html/template"
	"net/http"
)

type ErrorPageMsg struct {
	errorCode string
	errorMsg  string
}

func errHandler(w http.ResponseWriter, r *http.Request, err *ErrorPageMsg) {
	errTmp := template.Must(template.ParseFiles("template/index.html"))
	errTmp.Execute(w, err)
}

// function to validate the main
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// check for the path request
	if r.URL.Path != "/" {
		err := ErrorPageMsg{errorCode: "404", errorMsg: "PAGE NOT FOUND"}
		w.WriteHeader(http.StatusFound)
		errHandler(w, r, &err)
		return
	}
	if r.Method != "Get" {
		err := ErrorPageMsg{errorCode: "405", errorMsg: "METHOD NOT ALLOWED"}
		w.WriteHeader(http.StatusMethodNotAllowed)
		errHandler(w, r, &err)
		return
	}
	//validating the parsing of the main page
	main, err := template.ParseFiles("template/index.html")
	if err != nil {
		err := ErrorPageMsg{errorCode: "500", errorMsg: "INTERNAL SERVER ERROR"}
		errHandler(w, r, &err)
		return
	}
	mainTmp := template.Must(main, nil)
	mainTmp.Execute(w, nil)
}

// function to hand index.html in the server.
func Hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
