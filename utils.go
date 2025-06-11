package main

import (
	"fmt"
	"math/rand"
	"time"
)

const shortCodeLength = 6

// rune means int32 but is used to distinguish character values
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

const maxAttempts = 100

func GenerateUniqueShortCode() (string, error) { // returns string and error
	for i := 0; i < maxAttempts; i++ {
		code := GenerateShortCode()
		if _, exists := GetURL(code); !exists {
			return code, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique short code after %d attempts", maxAttempts)
}

// creates a random 6-character string
func GenerateShortCode() string {
	b := make([]rune, shortCodeLength) // array of length 6
	for i := range b {
		b[i] = letters[seededRand.Intn(len(letters))]
	}
	return string(b)
}
