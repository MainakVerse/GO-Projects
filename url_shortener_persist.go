package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const dbFile = "urls.json"

var urlStore map[string]string

// Load saved URLs from file
func loadDB() {
	file, err := os.ReadFile(dbFile)
	if err != nil {
		urlStore = make(map[string]string)
		return
	}
	json.Unmarshal(file, &urlStore)
}

// Save URLs to file
func saveDB() {
	data, _ := json.MarshalIndent(urlStore, "", "  ")
	os.WriteFile(dbFile, data, 0644)
}

// Generate a short random code
func generateShortCode(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:n]
}

func shorten(url string) string {
	short := generateShortCode(6)
	urlStore[short] = url
	saveDB()
	return short
}

func expand(code string) string {
	if val, ok := urlStore[code]; ok {
		return val
	}
	return "❌ No URL found for this code."
}

func main() {
	loadDB()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔗 Persistent URL Shortener")
	fmt.Println("----------------------------")

	for {
		fmt.Print("\nCommand (shorten/expand/list/exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		switch cmd {
		case "shorten":
			fmt.Print("Enter long URL: ")
			long, _ := reader.ReadString('\n')
			long = strings.TrimSpace(long)
			code := shorten(long)
			fmt.Printf("✅ Short URL: %s → %s\n", code, long)

		case "expand":
			fmt.Print("Enter short code: ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)
			fmt.Println("🔗 Original URL:", expand(code))

		case "list":
			if len(urlStore) == 0 {
				fmt.Println("ℹ️ No URLs stored yet.")
				continue
			}
			fmt.Println("📋 Stored URLs:")
			for k, v := range urlStore {
				fmt.Printf("%s → %s\n", k, v)
			}

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Unknown command.")
		}
	}
}
