package main

import (
	"fmt"
	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL required", http.StatusBadRequest)
		return
	}

	code, err := GenerateUniqueShortCode()
	if err != nil {
		http.Error(w, "Could not generate short code", http.StatusInternalServerError)
		return
	}

	SaveURL(code, url)
	AddToRecent(code, url)

	// Show the user their short URL
	fmt.Fprintf(w, "Short URL: http://localhost:8080/s/%s", code)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/s/")
	originalURL, exists := GetURL(code)
	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
