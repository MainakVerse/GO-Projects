package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Joke struct {
	Setup   string `json:"setup"`
	Punch   string `json:"punchline"`
}

func main() {
	for {
		fmt.Print("\nPress ENTER to get a random joke or type 'exit': ")
		var cmd string
		fmt.Scanln(&cmd)
		if cmd == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
		if err != nil {
			fmt.Println("❌ Failed to fetch joke:", err)
			continue
		}
		defer resp.Body.Close()

		var j Joke
		if err := json.NewDecoder(resp.Body).Decode(&j); err != nil {
			fmt.Println("❌ Error parsing joke:", err)
			continue
		}

		fmt.Printf("\n😂 %s\n🤣 %s\n", j.Setup, j.Punch)
	}
}
