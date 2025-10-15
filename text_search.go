package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter file path: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Print("Enter search term: ")
	searchTerm, _ := reader.ReadString('\n')
	searchTerm = strings.ToLower(strings.TrimSpace(searchTerm))

	fmt.Println("\n🔎 Matching lines:\n-----------------")

	scanner := bufio.NewScanner(file)
	lineNum := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), searchTerm) {
			fmt.Printf("[%d] %s\n", lineNum, line)
			found = true
		}
		lineNum++
	}

	if !found {
		fmt.Println("No matches found.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("⚠️ Error reading file:", err)
	}
}
