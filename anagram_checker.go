package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortString(s string) string {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔤 Anagram Checker")
	fmt.Println("-------------------")

	for {
		fmt.Print("\nEnter first word (or 'exit'): ")
		a, _ := reader.ReadString('\n')
		a = strings.TrimSpace(a)
		if strings.ToLower(a) == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		fmt.Print("Enter second word: ")
		b, _ := reader.ReadString('\n')
		b = strings.TrimSpace(b)

		if sortString(a) == sortString(b) {
			fmt.Println("✅ They are anagrams!")
		} else {
			fmt.Println("❌ Not anagrams.")
		}
	}
}
