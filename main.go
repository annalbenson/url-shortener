package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed static/*
var staticFiles embed.FS

var tmpl = template.Must(template.ParseFS(staticFiles, "static/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	recent := GetRecent()
	err := tmpl.Execute(w, recent)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("/shorten", ShortenFormHandler)
	http.HandleFunc("/s/", RedirectHandler)

	log.Printf("Starting server")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
