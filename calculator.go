package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readLineFloat(r *bufio.Reader, prompt string) (float64, error) {
	fmt.Print(prompt)
	s, err := r.ReadString('\n')
	if err != nil {
		return 0, err
	}
	s = strings.TrimSpace(s)
	return strconv.ParseFloat(s, 64)
}

func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nChoose operation (+, -, *, /) or type exit: ")
		opRaw, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Read error:", err)
			return
		}
		op := strings.TrimSpace(opRaw)
		if op == "exit" {
			fmt.Println("👋 Goodbye!")
			return
		}

		a, err := readLineFloat(r, "Enter first number: ")
		if err != nil {
			fmt.Println("Invalid number:", err)
			continue
		}
		b, err := readLineFloat(r, "Enter second number: ")
		if err != nil {
			fmt.Println("Invalid number:", err)
			continue
		}

		switch op {
		case "+":
			fmt.Printf("Result: %.6g\n", a+b)
		case "-":
			fmt.Printf("Result: %.6g\n", a-b)
		case "*":
			fmt.Printf("Result: %.6g\n", a*b)
		case "/":
			if b == 0 {
				fmt.Println("❌ Cannot divide by zero")
			} else {
				fmt.Printf("Result: %.6g\n", a/b)
			}
		default:
			fmt.Println("❌ Invalid operation")
		}
	}
}
