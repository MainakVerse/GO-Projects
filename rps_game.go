package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	choices := []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())

	fmt.Println("✊📄✂️  Rock Paper Scissors Game")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("\nEnter your choice (rock/paper/scissors or exit): ")
		var player string
		fmt.Scanln(&player)
		player = strings.ToLower(player)

		if player == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		valid := false
		for _, c := range choices {
			if player == c {
				valid = true
				break
			}
		}
		if !valid {
			fmt.Println("❌ Invalid choice, try again.")
			continue
		}

		computer := choices[rand.Intn(3)]
		fmt.Printf("🤖 Computer chose: %s\n", computer)

		switch {
		case player == computer:
			fmt.Println("⚖️  It's a tie!")
		case (player == "rock" && computer == "scissors") ||
			(player == "scissors" && computer == "paper") ||
			(player == "paper" && computer == "rock"):
			fmt.Println("🏆 You win!")
		default:
			fmt.Println("💀 You lose!")
		}
	}
}
