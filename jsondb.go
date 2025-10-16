package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dbFile = "data.json"

var db map[string]string

func loadDB() {
	file, err := os.ReadFile(dbFile)
	if err != nil {
		db = make(map[string]string)
		return
	}
	json.Unmarshal(file, &db)
}

func saveDB() {
	data, _ := json.MarshalIndent(db, "", "  ")
	os.WriteFile(dbFile, data, 0644)
}

func main() {
	loadDB()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("💾 JSON-based Mini Database")
	fmt.Println("------------------------------")

	for {
		fmt.Print("\nCommand (set/get/del/show/exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		switch cmd {
		case "set":
			fmt.Print("Key: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			fmt.Print("Value: ")
			val, _ := reader.ReadString('\n')
			val = strings.TrimSpace(val)
			db[key] = val
			saveDB()
			fmt.Println("✅ Saved.")

		case "get":
			fmt.Print("Key: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			if val, ok := db[key]; ok {
				fmt.Println("🔍 Value:", val)
			} else {
				fmt.Println("❌ Key not found.")
			}

		case "del":
			fmt.Print("Key: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			delete(db, key)
			saveDB()
			fmt.Println("🗑️ Deleted (if existed).")

		case "show":
			if len(db) == 0 {
				fmt.Println("ℹ️ Database is empty.")
				continue
			}
			fmt.Println("📋 Records:")
			for k, v := range db {
				fmt.Printf("%s = %s\n", k, v)
			}

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Unknown command.")
		}
	}
}
