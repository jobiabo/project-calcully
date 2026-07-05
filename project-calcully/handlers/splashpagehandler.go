package handler

import (
	"html/template"
	"net/http"
)

type PageOutput struct {
	Input  string
	Result string
}

// var powers = "㎡"
// var tmpl = template.Must(template.ParseFiles("template/claude2.html"))
// var tmpl = template.Must(template.ParseFiles("template/gemini6.html"))
var templ = template.Must(template.ParseFiles("template/splashpage.html"))

func SplashPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	templ.Execute(w, nil)
	// http.Redirect(w, r, "/splash/home", http.StatusMovedPermanently)

}
