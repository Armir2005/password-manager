package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}

var account []Account
var isLoggedIn bool = false

func createAccount(username, password string) {
	loadAccount()

	for _, acc := range account {
		if acc.Username == username {
			fmt.Println("Username already exists. Please choose a different username.")
			return
		}
	}

	salt := generateSalt()
	encryptedUsername := encrypt(username, deriveKey(password, salt))
	encryptedPassword := encrypt(password, deriveKey(password, salt))
	account = append(account, Account{encryptedUsername, encryptedPassword, salt})
	saveAccount()
	fmt.Println("Account saved.")
}

func saveAccount() {
	file, err := os.Create("account.json")
	if err != nil {
		fmt.Println("Error creating account file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(account)
}

func loadAccount() {
	file, err := os.Open("account.json")
	if err != nil {
		fmt.Println("Error loading accounts:", err)
		return
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&account); err != nil {
		fmt.Println("Error decoding account data:", err)
	}
}

func loginAccount(username, password string) bool {
	for _, acc := range account {
		key := deriveKey(password, acc.Salt)
		decryptedUsername := decrypt(acc.Username, key)
		if decryptedUsername == username {
			decryptedPassword := decrypt(acc.Password, key)
			if decryptedPassword == password {
				isLoggedIn = true
				sessionKey = key
				fmt.Println("Login successful.")
				return true
			}
		}
	}
	fmt.Println("Invalid username or password.")
	return false
}
