package main

import (
	"bufio"
	"fmt"
	"os"
)

var todos []string

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n1. Add | 2. List | 3. Delete | 4. Exit")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			todos = append(todos, scanner.Text())

		case "2":
			fmt.Println("🗒️  Your Tasks:")
			for i, t := range todos {
				fmt.Printf("%d. %s\n", i+1, t)
			}

		case "3":
			fmt.Print("Enter task number: ")
			scanner.Scan()
			var n int
			fmt.Sscan(scanner.Text(), &n)
			if n > 0 && n <= len(todos) {
				todos = append(todos[:n-1], todos[n:]...)
				fmt.Println("✅ Task deleted")
			} else {
				fmt.Println("❌ Invalid number")
			}

		case "4":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid choice!")
		}
	}
}
