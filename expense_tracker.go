package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const dbFile = "expenses.json"

type Expense struct {
	Item     string  `json:"item"`
	Amount   float64 `json:"amount"`
	Category string  `json:"category"`
	Date     string  `json:"date"`
}

var expenses []Expense

func loadExpenses() {
	data, err := os.ReadFile(dbFile)
	if err != nil {
		expenses = []Expense{}
		return
	}
	json.Unmarshal(data, &expenses)
}

func saveExpenses() {
	data, _ := json.MarshalIndent(expenses, "", "  ")
	os.WriteFile(dbFile, data, 0644)
}

func addExpense(item string, amount float64, category string) {
	e := Expense{item, amount, category, time.Now().Format("2006-01-02")}
	expenses = append(expenses, e)
	saveExpenses()
	fmt.Println("✅ Expense added successfully!")
}

func viewExpenses() {
	if len(expenses) == 0 {
		fmt.Println("ℹ️ No expenses recorded yet.")
		return
	}

	total := 0.0
	fmt.Println("\n📊 Expense List:")
	for _, e := range expenses {
		fmt.Printf("%-10s | ₹%.2f | %s | %s\n", e.Item, e.Amount, e.Category, e.Date)
		total += e.Amount
	}
	fmt.Printf("\n💵 Total Spent: ₹%.2f\n", total)
}

func main() {
	loadExpenses()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("💰 Simple Expense Tracker")
	fmt.Println("---------------------------")

	for {
		fmt.Print("\nCommand (add/view/exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		switch cmd {
		case "add":
			fmt.Print("Item name: ")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSpace(item)

			fmt.Print("Amount: ")
			amtStr, _ := reader.ReadString('\n')
			amtStr = strings.TrimSpace(amtStr)
			amount, err := strconv.ParseFloat(amtStr, 64)
			if err != nil {
				fmt.Println("❌ Invalid amount.")
				continue
			}

			fmt.Print("Category: ")
			category, _ := reader.ReadString('\n')
			category = strings.TrimSpace(category)

			addExpense(item, amount, category)

		case "view":
			viewExpenses()

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Unknown command.")
		}
	}
}
