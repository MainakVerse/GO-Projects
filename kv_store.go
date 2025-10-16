package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var store = make(map[string]string)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("💾 Simple Key-Value Store")
	fmt.Println("----------------------------")

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
			store[key] = val
			fmt.Println("✅ Saved.")

		case "get":
			fmt.Print("Key: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			if val, ok := store[key]; ok {
				fmt.Println("🔍 Value:", val)
			} else {
				fmt.Println("❌ Key not found.")
			}

		case "del":
			fmt.Print("Key: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			delete(store, key)
			fmt.Println("🗑️ Deleted (if existed).")

		case "show":
			if len(store) == 0 {
				fmt.Println("ℹ️ Store is empty.")
				continue
			}
			fmt.Println("📦 Current Data:")
			for k, v := range store {
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
