package main

import "sync"

var (
	urlStore = make(map[string]string)
	mu       sync.RWMutex
)

// store shortCode --> originalURL pair
func SaveURL(code string, original string) {
	mu.Lock()
	defer mu.Unlock()
	urlStore[code] = original
}

func GetURL(code string) (string, bool) { // returns string and bool
	mu.RLock()
	defer mu.RUnlock()
	original, exists := urlStore[code]
	return original, exists
}

type Shortened struct {
	Code string
	URL  string
}

var recentList []Shortened

const maxRecent = 10

func AddToRecent(code, original string) {
	mu.Lock()
	defer mu.Unlock()

	recentList = append([]Shortened{{code, original}}, recentList...)
	if len(recentList) > maxRecent {
		recentList = recentList[:maxRecent]
	}
}

func GetRecent() []Shortened {
	mu.RLock()
	defer mu.RUnlock()
	return recentList
}
