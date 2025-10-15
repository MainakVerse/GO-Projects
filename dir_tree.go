package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func printTree(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("❌ Error reading:", err)
		return
	}

	for i, entry := range entries {
		connector := "├──"
		if i == len(entries)-1 {
			connector = "└──"
		}
		fmt.Println(prefix + connector + " " + entry.Name())

		if entry.IsDir() {
			newPrefix := prefix
			if i == len(entries)-1 {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			printTree(filepath.Join(path, entry.Name()), newPrefix)
		}
	}
}

func main() {
	fmt.Print("Enter directory path: ")
	var dir string
	fmt.Scanln(&dir)

	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println(info.Name())
	printTree(dir, "")
}
