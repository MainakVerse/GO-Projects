package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter CSV file path: ")
	var path string
	fmt.Scanln(&path)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("❌ Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("❌ Error reading CSV:", err)
		return
	}

	if len(records) < 2 {
		fmt.Println("❌ CSV must have headers and at least one row")
		return
	}

	headers := records[0]
	var result []map[string]string

	for _, row := range records[1:] {
		item := map[string]string{}
		for i, val := range row {
			item[headers[i]] = val
		}
		result = append(result, item)
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("❌ Error converting to JSON:", err)
		return
	}

	fmt.Print("Enter output JSON file name (or leave blank to print): ")
	var out string
	fmt.Scanln(&out)

	if out == "" {
		fmt.Println("\n✅ JSON Output:\n---------------")
		fmt.Println(string(jsonData))
	} else {
		if err := os.WriteFile(out, jsonData, 0644); err != nil {
			fmt.Println("❌ Error writing JSON:", err)
			return
		}
		fmt.Println("✅ JSON saved to:", out)
	}
}
