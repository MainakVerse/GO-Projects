package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rates := map[string]float64{
		"USD": 1.0,     // Base currency
		"EUR": 0.92,
		"INR": 83.1,
		"GBP": 0.78,
		"JPY": 151.5,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nSupported: USD, EUR, INR, GBP, JPY")
		fmt.Print("From currency: ")
		from, _ := reader.ReadString('\n')
		from = strings.ToUpper(strings.TrimSpace(from))

		fmt.Print("To currency: ")
		to, _ := reader.ReadString('\n')
		to = strings.ToUpper(strings.TrimSpace(to))

		if _, ok := rates[from]; !ok {
			fmt.Println("❌ Invalid FROM currency")
			continue
		}
		if _, ok := rates[to]; !ok {
			fmt.Println("❌ Invalid TO currency")
			continue
		}

		fmt.Print("Amount: ")
		amtStr, _ := reader.ReadString('\n')
		amtStr = strings.TrimSpace(amtStr)
		amount, err := strconv.ParseFloat(amtStr, 64)
		if err != nil {
			fmt.Println("❌ Invalid amount")
			continue
		}

		usd := amount / rates[from]       // Convert to USD
		converted := usd * rates[to]      // Convert to target
		fmt.Printf("%.2f %s = %.2f %s\n", amount, from, converted, to)

		fmt.Print("Convert again? (y/n): ")
		ans, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(ans)) != "y" {
			fmt.Println("👋 Goodbye!")
			break
		}
	}
}
