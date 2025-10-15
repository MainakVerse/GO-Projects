package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter folder path: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("❌ Error reading directory:", err)
		return
	}

	fmt.Printf("Found %d files.\n", len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fmt.Printf("\nCurrent name: %s\nNew name (or press Enter to skip): ", file.Name())
		newName, _ := reader.ReadString('\n')
		newName = strings.TrimSpace(newName)

		if newName == "" {
			continue
		}

		oldPath := filepath.Join(path, file.Name())
		newPath := filepath.Join(path, newName)
		if err := os.Rename(oldPath, newPath); err != nil {
			fmt.Println("❌ Error renaming:", err)
		} else {
			fmt.Println("✅ Renamed to:", newName)
		}
	}

	fmt.Println("\n🎉 All done!")
}
