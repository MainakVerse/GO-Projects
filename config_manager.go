package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func writeFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}

func main() {
	fmt.Print("Enter config file path (.json or .yaml): ")
	var path string
	fmt.Scanln(&path)

	data, err := readFile(path)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return
	}

	config := make(map[string]interface{})
	if strings.HasSuffix(path, ".json") {
		json.Unmarshal(data, &config)
	} else if strings.HasSuffix(path, ".yaml") || strings.HasSuffix(path, ".yml") {
		yaml.Unmarshal(data, &config)
	} else {
		fmt.Println("❌ Unsupported file format.")
		return
	}

	fmt.Println("\n📄 Current Config:")
	for k, v := range config {
		fmt.Printf("%s = %v\n", k, v)
	}

	fmt.Print("\nDo you want to update a key? (y/n): ")
	var choice string
	fmt.Scanln(&choice)
	if strings.ToLower(choice) != "y" {
		fmt.Println("👋 Exiting.")
		return
	}

	fmt.Print("Enter key: ")
	var key string
	fmt.Scanln(&key)
	fmt.Print("Enter new value: ")
	var value string
	fmt.Scanln(&value)
	config[key] = value

	var output []byte
	if strings.HasSuffix(path, ".json") {
		output, _ = json.MarshalIndent(config, "", "  ")
	} else {
		output, _ = yaml.Marshal(config)
	}

	if err := writeFile(path, output); err != nil {
		fmt.Println("❌ Error saving file:", err)
		return
	}

	fmt.Println("✅ Configuration updated successfully!")
}
