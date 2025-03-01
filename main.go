package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Password Manager:")
	fmt.Println("Type 'help', to show available commands.")

	var loggedInUser string
	var specialChars bool

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  create [Username] [Password]                                      - Create an account")
			fmt.Println("  login [Username] [Password]                                       - Login to your account")
			fmt.Println("  logout                                                            - Logout from your account")
			fmt.Println("  add [Service] [Username] [Password]                               - Add a password")
			fmt.Println("  generate [Service] [Username] [Length] [Special Characters (y/n)] - Generates a password")
			fmt.Println("  list                                                              - List all passwords")
			fmt.Println("  clear                                                             - Clear the terminal")
			fmt.Println("  exit                                                              - Exit the program")
		case "create":
			if len(parts) < 3 {
				fmt.Println("Usage: create [Username] [Password]")
				continue
			}
			username := parts[1]
			password := parts[2]
			createAccount(username, password)
		case "login":
			if len(parts) < 3 {
				fmt.Println("Usage: login [Username] [Password]")
				continue
			}
			username := parts[1]
			password := parts[2]
			loadAccount()
			if loginAccount(username, password) {
				loggedInUser = username
			}
		case "logout":
			if !isLoggedIn {
				fmt.Println("You are not logged in.")
				continue
			}
			loggedInUser = ""
			isLoggedIn = false
			sessionKey = nil
			fmt.Println("Logged out successfully.")
		case "add":
			if len(parts) < 4 {
				fmt.Println("Usage: add [Service] [Username] [Password]")
				continue
			}
			service := parts[1]
			username := parts[2]
			password := parts[3]
			addPassword(loggedInUser, service, username, password)
		case "generate":
			if len(parts) < 5 {
				fmt.Println("Usage: generate [Service] [Username] [Length] [Special Characters (y/n)]")
				continue
			}
			service := parts[1]
			username := parts[2]
			length, err := strconv.Atoi(parts[3])
			if err != nil {
				fmt.Println("Length must be a number.")
				continue
			}
			special := parts[4]
			if special == "y" {
				specialChars = true
			} else if special == "n" {
				specialChars = false
			} else {
				fmt.Println("[Special Characters] must be y or n.")
				continue
			}
			addPassword(loggedInUser, service, username, generatePassword(length, specialChars))
		case "list":
			listPasswords(loggedInUser)
		case "clear":
			fmt.Print("\033[H\033[2J")
			continue
		case "exit":
			fmt.Print("Exiting...\n")
			return
		default:
			fmt.Println("Invalid command. Type 'help' to get help.")
		}
	}
}
