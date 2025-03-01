package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PasswordEntry struct {
	Service  string `json:"service"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var passwords []PasswordEntry

func savePasswords(username string) {
	file, _ := os.Create(username + "_passwords.json")
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(passwords)
}

func loadPasswords(username string) {
	file, err := os.Open(username + "_passwords.json")
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&passwords)
}

func addPassword(username, service, user, password string) {
	if !isLoggedIn {
		fmt.Println("You must be logged in to add a password.")
		return
	}
	encryptedService := encrypt(service, sessionKey)
	encryptedUser := encrypt(user, sessionKey)
	encryptedPassword := encrypt(password, sessionKey)
	passwords = append(passwords, PasswordEntry{encryptedService, encryptedUser, encryptedPassword})
	savePasswords(username)
	fmt.Println("Password saved.")
}

func listPasswords(username string) {
	if !isLoggedIn {
		fmt.Println("You must be logged in to see the passwords.")
		return
	}
	loadPasswords(username)
	for _, entry := range passwords {
		fmt.Printf("Service: %s, Username: %s, Password: %s\n", decrypt(entry.Service, sessionKey), decrypt(entry.Username, sessionKey), decrypt(entry.Password, sessionKey))
	}
}
