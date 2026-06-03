package server

import (
	"net/http"
	"text/template"
)

type ErrorPage struct {
	errcode string
	errmsg  string
}

var Err = *ErrorPage

func errHandler(w http.ResponseWriter, r *http.Request) {
	errTmp := template.Must(template.ParseFiles("template/error.html"))
	errTmp.Execute(w, Err)
}
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := ErrorPage{errcode: "404", errmsg: "PAGE NOT FOUND"}
		w.WriteHeader(http.StatusNotFound)
		errHandler(w, r, &Err)
		return
	}
}
