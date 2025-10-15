package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func main() {
	fmt.Print("Enter folder path: ")
	var dir string
	fmt.Scanln(&dir)

	hashMap := make(map[string][]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		hash, err := hashFile(path)
		if err == nil {
			hashMap[hash] = append(hashMap[hash], path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("❌ Error scanning directory:", err)
		return
	}

	fmt.Println("\n🔍 Duplicate Files Found:")
	found := false
	for _, files := range hashMap {
		if len(files) > 1 {
			found = true
			fmt.Println("-----------------------")
			for _, f := range files {
				fmt.Println(f)
			}
		}
	}

	if !found {
		fmt.Println("✅ No duplicates found.")
	}
}
