package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	fmt.Print("Enter file path: ")
	var path string
	fmt.Scanln(&path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println("\n📄 File Metadata")
	fmt.Println("--------------------------")
	fmt.Printf("File Name   : %s\n", info.Name())
	fmt.Printf("Full Path   : %s\n", filepath.Join(path))
	fmt.Printf("Size        : %.2f KB\n", float64(info.Size())/1024)
	fmt.Printf("Permissions : %s\n", info.Mode())
	fmt.Printf("Modified At : %s\n", info.ModTime().Format(time.RFC1123))
	fmt.Printf("Is Directory: %v\n", info.IsDir())

	// Optional: show file type (by extension)
	ext := filepath.Ext(path)
	if ext != "" {
		fmt.Printf("Extension   : %s\n", ext)
	}
}
