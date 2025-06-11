package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("static/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	recent := GetRecent()
	tmpl.Execute(w, recent)
}

func main() {
	http.HandleFunc("/", HomeHandler)

	http.HandleFunc("/shorten", ShortenFormHandler)
	http.HandleFunc("/s/", RedirectHandler)

	http.ListenAndServe(":8080", nil)
}
