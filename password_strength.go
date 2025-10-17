package main

import (
	"fmt"
	"unicode"
)

func checkStrength(password string) string {
	var length, upper, lower, digit, special bool

	if len(password) >= 8 {
		length = true
	}
	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			upper = true
		case unicode.IsLower(ch):
			lower = true
		case unicode.IsDigit(ch):
			digit = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			special = true
		}
	}

	score := 0
	if length {
		score++
	}
	if upper {
		score++
	}
	if lower {
		score++
	}
	if digit {
		score++
	}
	if special {
		score++
	}

	switch score {
	case 5:
		return "💪 Very Strong"
	case 4:
		return "✅ Strong"
	case 3:
		return "⚠️ Medium"
	case 2:
		return "❌ Weak"
	default:
		return "💀 Very Weak"
	}
}

func main() {
	for {
		var pw string
		fmt.Print("\nEnter password (or 'exit' to quit): ")
		fmt.Scanln(&pw)
		if pw == "exit" {
			fmt.Println("👋 Goodbye!")
			break
		}
		fmt.Println("🔎 Strength:", checkStrength(pw))
	}
}
