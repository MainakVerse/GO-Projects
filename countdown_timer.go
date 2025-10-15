package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter countdown time in seconds (or 0 to exit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		seconds, err := strconv.Atoi(input)

		if err != nil || seconds < 0 {
			fmt.Println("❌ Invalid number")
			continue
		}
		if seconds == 0 {
			fmt.Println("👋 Goodbye!")
			return
		}

		fmt.Println("⏳ Countdown started...")
		for i := seconds; i > 0; i-- {
			fmt.Printf("\rTime left: %d seconds", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Print("\r⏰ Time’s up!                \n")

		fmt.Print("Start another countdown? (y/n): ")
		ans, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(ans)) != "y" {
			fmt.Println("👋 Goodbye!")
			break
		}
	}
}
