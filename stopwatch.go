package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var start time.Time
	var elapsed time.Duration
	running := false

	fmt.Println("🕒 Simple Stopwatch CLI")
	fmt.Println("Commands: start | stop | reset | exit")

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(strings.ToLower(input))

		switch cmd {
		case "start":
			if !running {
				start = time.Now().Add(-elapsed)
				running = true
				fmt.Println("▶️ Started!")
			} else {
				fmt.Println("⏸ Already running.")
			}

		case "stop":
			if running {
				elapsed = time.Since(start)
				running = false
				fmt.Printf("⏹ Stopped. Time: %v\n", elapsed.Round(time.Millisecond))
			} else {
				fmt.Println("❌ Not running.")
			}

		case "reset":
			start = time.Now()
			elapsed = 0
			if running {
				fmt.Println("🔄 Reset and running.")
			} else {
				fmt.Println("🔄 Reset to 0.")
			}

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			if running {
				fmt.Printf("⏱  Elapsed: %v\n", time.Since(start).Round(time.Millisecond))
			} else {
				fmt.Printf("⏱  Last recorded: %v\n", elapsed.Round(time.Millisecond))
			}
		}
	}
}
