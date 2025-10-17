package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	for {
		fmt.Print("\nPress ENTER for a random quote or type 'exit': ")
		var cmd string
		fmt.Scanln(&cmd)
		if cmd == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		resp, err := http.Get("https://api.quotable.io/random")
		if err != nil {
			fmt.Println("❌ Failed to fetch quote:", err)
			continue
		}
		defer resp.Body.Close()

		var q Quote
		if err := json.NewDecoder(resp.Body).Decode(&q); err != nil {
			fmt.Println("❌ Error decoding quote:", err)
			continue
		}

		fmt.Printf("\n💬 \"%s\"\n— %s\n", q.Content, q.Author)
	}
}
