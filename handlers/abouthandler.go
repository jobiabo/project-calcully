package handler

import (
	"net/http"
	"text/template"
)

var templa = template.Must(template.ParseFiles("template/aboutpage.html"))

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	templa.Execute(w, nil)

}
