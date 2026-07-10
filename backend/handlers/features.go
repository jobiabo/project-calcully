package handler

import (
	"html/template"
	"net/http"
)

var templf = template.Must(template.ParseFiles("template/features.html"))

func FeatureHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/features" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	templf.Execute(w, nil)
}
