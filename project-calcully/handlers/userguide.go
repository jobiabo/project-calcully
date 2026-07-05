package handler

import (
	"html/template"
	"net/http"
)

var templu = template.Must(template.ParseFiles("template/userguide.html"))

func UserGuide(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/userguide" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	templu.Execute(w, nil)
}
