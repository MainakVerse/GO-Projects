package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func main() {
	quiz := []Question{
		{"What is the capital of France?", "paris"},
		{"What is 5 + 3?", "8"},
		{"Who wrote 'Go' language?", "google"},
		{"What is the color of the sky?", "blue"},
		{"What is 10 / 2?", "5"},
	}

	reader := bufio.NewReader(os.Stdin)
	score := 0

	fmt.Println("🧩 Welcome to the Go Quiz!")
	fmt.Println("-------------------------")

	for i, q := range quiz {
		fmt.Printf("\nQ%d: %s\n> ", i+1, q.question)
		ans, _ := reader.ReadString('\n')
		ans = strings.TrimSpace(strings.ToLower(ans))

		if ans == q.answer {
			fmt.Println("✅ Correct!")
			score++
		} else {
			fmt.Printf("❌ Wrong! Correct answer: %s\n", q.answer)
		}
	}

	fmt.Printf("\n🏁 Quiz Over! Your Score: %d/%d\n", score, len(quiz))
	fmt.Println("🎉 Thanks for playing!")
}
