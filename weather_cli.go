package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	for {
		fmt.Print("\nEnter city name (or 'exit' to quit): ")
		var city string
		fmt.Scanln(&city)
		if strings.ToLower(city) == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		url := fmt.Sprintf("https://wttr.in/%s?format=3", city)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("❌ Network error:", err)
			continue
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		fmt.Println("🌤️  Weather:", string(data))
	}
}
