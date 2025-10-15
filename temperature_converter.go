package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTemperature Converter:")
		fmt.Println("1. Celsius → Fahrenheit")
		fmt.Println("2. Fahrenheit → Celsius")
		fmt.Println("3. Celsius → Kelvin")
		fmt.Println("4. Kelvin → Celsius")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "5" {
			fmt.Println("👋 Goodbye!")
			return
		}

		fmt.Print("Enter value: ")
		valStr, _ := reader.ReadString('\n')
		valStr = strings.TrimSpace(valStr)
		value, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			fmt.Println("❌ Invalid number")
			continue
		}

		switch choice {
		case "1":
			fmt.Printf("%.2f°C = %.2f°F\n", value, value*9/5+32)
		case "2":
			fmt.Printf("%.2f°F = %.2f°C\n", value, (value-32)*5/9)
		case "3":
			fmt.Printf("%.2f°C = %.2fK\n", value, value+273.15)
		case "4":
			fmt.Printf("%.2fK = %.2f°C\n", value, value-273.15)
		default:
			fmt.Println("❌ Invalid option")
		}
	}
}
