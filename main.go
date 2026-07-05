package main

import (
	handler "calcully/handlers"
	"log"
	"net/http"
)

const powers = "㎡"

func main() {

	fs := http.FileServer(http.Dir("static")) // ✅ matches where IMG_9139.JPG actually is
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler.SplashPageHandler)
	http.HandleFunc("/home", handler.HomePageHandler)
	http.HandleFunc("/calculator", handler.CalcXPageHandler)
	http.HandleFunc("/userguide", handler.UserGuide)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/features", handler.FeatureHandlers)

	log.Println("server running at http://localhost:9000")
	// http.ListenAndServe(":9000", mux)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

}
