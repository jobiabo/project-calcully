package handler

import (
	"html/template"
	"net/http"
)

var templh = template.Must(template.ParseFiles("template/homepage.html"))

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	templh.Execute(w, nil)
}
