package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Print("Enter folder path: ")
	var path string
	fmt.Scanln(&path)

	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("⚠️ Error accessing:", err)
			return nil
		}
		if !info.IsDir() {
			sizeKB := float64(info.Size()) / 1024
			fmt.Printf("%-50s %.2f KB\n", info.Name(), sizeKB)
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("\n✅ Scan complete!")
	}
}
