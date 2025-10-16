package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"
)

var urlStore = make(map[string]string)

func generateShortCode(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:n]
}

func shorten(url string) string {
	short := generateShortCode(6)
	urlStore[short] = url
	return short
}

func expand(short string) string {
	if long, ok := urlStore[short]; ok {
		return long
	}
	return "❌ No such short URL found."
}

func main() {
	for {
		fmt.Println("\n1. Shorten URL\n2. Expand short URL\n3. Show all\n4. Exit")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter long URL: ")
			var url string
			fmt.Scanln(&url)
			short := shorten(url)
			fmt.Printf("✅ Short URL: %s -> %s\n", short, url)

		case 2:
			fmt.Print("Enter short code: ")
			var code string
			fmt.Scanln(&code)
			fmt.Println("🔗 Original URL:", expand(strings.TrimSpace(code)))

		case 3:
			fmt.Println("\nStored URLs:")
			for k, v := range urlStore {
				fmt.Printf("%s -> %s\n", k, v)
			}

		case 4:
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice.")
		}
	}
}
