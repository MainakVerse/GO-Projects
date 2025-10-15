package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter log file path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Print("Enter filter keyword (e.g. ERROR, INFO, WARN or text): ")
	filter, _ := reader.ReadString('\n')
	filter = strings.ToLower(strings.TrimSpace(filter))

	fmt.Println("\n🔎 Matching log entries:\n------------------------")

	scanner := bufio.NewScanner(file)
	lineNum := 1
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(strings.ToLower(line), filter) {
			fmt.Printf("[%d] %s\n", lineNum, line)
			found = true
		}
		lineNum++
	}

	if !found {
		fmt.Println("No matching entries found.")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("⚠️ Error reading file:", err)
	}
}
