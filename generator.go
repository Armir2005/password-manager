package main

import (
	"math/rand"
	"time"
)

func generatePassword(length int, specialChars bool) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	const special = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"

	var charset string
	if specialChars {
		charset = letters + special
	} else {
		charset = letters
	}

	rand.Seed(time.Now().UnixNano())
	generatedPassword := make([]byte, length)
	for i := range generatedPassword {
		generatedPassword[i] = charset[rand.Intn(len(charset))]
	}

	return string(generatedPassword)
}
