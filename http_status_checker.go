package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		fmt.Print("\nEnter URL to check (or 'exit' to quit): ")
		var url string
		fmt.Scanln(&url)

		if url == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		// Add scheme if missing
		if !(len(url) >= 4 && (url[:4] == "http")) {
			url = "https://" + url
		}

		client := http.Client{Timeout: 5 * time.Second}
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("❌ Error:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Printf("✅ %s -> %d %s\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}
