package server

import (
	"html/template"
	"net/http"
	"strings"
)

type ErrorPageMsg struct {
	errorCode string
	errorMsg  string
}

type resPageData struct {
	str  string
	font string
	res  string
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		errHandler(w, r, &err)
		return
	}
	mainTmp := template.Must(main, nil)
	mainTmp.Execute(w, nil)
}

// function for reshandler
func ResHandler(w http.ResponseWriter, r *http.Request) {
	// checking for parsing of the form
	err := r.ParseForm()
	if err != nil {
		err := ErrorPageMsg{errorCode: "500", errorMsg: "INTERNAL SERVER ERROR"}
		w.WriteHeader(http.StatusInternalServerError)
		errHandler(w, r, &err)
		return
	}
	str := r.PostFormValue("input-txt")
	validatestr := strings.ReplaceAll(str, "\r\n", "")

	for _, letts := range validatestr {
		if letts < 32 || letts > 126 {
			err := ErrorPageMsg{errorCode: "400", errorMsg: "INVALID INPUT"}
			w.WriteHeader(http.StatusNotAcceptable)
			errHandler(w, r, &err)
			return
		}
	}
	font := r.PostFormValue("banners")

	if font != "standard" && font != "shadow" && font != "thinkertoy" {
		err := ErrorPageMsg{errorCode: "404", errorMsg: "BANNER NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errHandler(w, r, &err)
		return
	}
	
}

// function to hand index.html in the server.
func Hello(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/index.html")
}
