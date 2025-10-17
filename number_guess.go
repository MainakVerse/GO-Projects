package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1 // random number between 1–100
	var guess int
	attempts := 0

	fmt.Println("🎯 Number Guessing Game")
	fmt.Println("-----------------------")
	fmt.Println("I'm thinking of a number between 1 and 100...")

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secret {
			fmt.Println("📉 Too low!")
		} else if guess > secret {
			fmt.Println("📈 Too high!")
		} else {
			fmt.Printf("🎉 Correct! The number was %d (in %d attempts)\n", secret, attempts)
			break
		}
	}
}
