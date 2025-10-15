package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file path (or leave blank for manual input): ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	var text string
	if path == "" {
		fmt.Println("Type text (end with an empty line):")
		for {
			line, _ := reader.ReadString('\n')
			if strings.TrimSpace(line) == "" {
				break
			}
			text += line
		}
	} else {
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("❌ Error reading file:", err)
			return
		}
		text = string(data)
	}

	lines := strings.Count(text, "\n")
	words := len(strings.Fields(text))
	chars := len([]rune(text))

	fmt.Println("\n📊 Word Count Report")
	fmt.Println("----------------------")
	fmt.Printf("Lines: %d\n", lines)
	fmt.Printf("Words: %d\n", words)
	fmt.Printf("Characters: %d\n", chars)
	fmt.Println("----------------------")
}
