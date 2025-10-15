package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎲 Dice Roller CLI — Roll virtual dice anytime!")

	for {
		fmt.Print("\nEnter number of dice (1-6) or 'exit': ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		var n int
		_, err := fmt.Sscan(input, &n)
		if err != nil || n <= 0 || n > 6 {
			fmt.Println("❌ Invalid number (1–6 only)")
			continue
		}

		fmt.Print("Rolling")
		for i := 0; i < 3; i++ {
			fmt.Print(".")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println()

		fmt.Print("Results: ")
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", rand.Intn(6)+1)
		}
		fmt.Println()

		fmt.Print("Roll again? (y/n): ")
		ans, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(ans)) != "y" {
			fmt.Println("👋 Goodbye!")
			break
		}
	}
}
