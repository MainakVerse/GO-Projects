package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dbFile = "users.json"

var users map[string]string

func loadUsers() {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		users = make(map[string]string)
		return
	}
	json.Unmarshal(data, &users)
}

func saveUsers() {
	data, _ := json.MarshalIndent(users, "", "  ")
	os.WriteFile(dbFile, data, 0644)
}

func hashPassword(pw string) string {
	h := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(h[:])
}

func main() {
	loadUsers()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔐 Simple User Credentials Manager")
	fmt.Println("-----------------------------------")

	for {
		fmt.Print("\nCommand (register/login/list/exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		switch cmd {
		case "register":
			fmt.Print("Enter username: ")
			user, _ := reader.ReadString('\n')
			user = strings.TrimSpace(user)
			if _, exists := users[user]; exists {
				fmt.Println("⚠️  Username already exists.")
				continue
			}

			fmt.Print("Enter password: ")
			pw, _ := reader.ReadString('\n')
			pw = strings.TrimSpace(pw)
			users[user] = hashPassword(pw)
			saveUsers()
			fmt.Println("✅ User registered successfully!")

		case "login":
			fmt.Print("Username: ")
			user, _ := reader.ReadString('\n')
			user = strings.TrimSpace(user)
			fmt.Print("Password: ")
			pw, _ := reader.ReadString('\n')
			pw = strings.TrimSpace(pw)

			if hashPassword(pw) == users[user] {
				fmt.Println("✅ Login successful!")
			} else {
				fmt.Println("❌ Invalid credentials.")
			}

		case "list":
			fmt.Println("\n📋 Registered Users:")
			for user := range users {
				fmt.Println("-", user)
			}

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid command.")
		}
	}
}
