package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zipFolder(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		relPath, _ := filepath.Rel(source, path)
		f, err := archive.Create(relPath)
		if err != nil {
			return err
		}
		fs, err := os.Open(path)
		if err != nil {
			return err
		}
		defer fs.Close()
		_, err = io.Copy(f, fs)
		return err
	})
	return nil
}

func unzip(source, target string) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		path := filepath.Join(target, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, os.ModePerm)
			continue
		}
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		out, err := os.Create(path)
		if err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		io.Copy(out, rc)
		out.Close()
		rc.Close()
	}
	return nil
}

func main() {
	var choice, src, dst string
	fmt.Println("1. Zip folder\n2. Unzip file")
	fmt.Print("Choose option: ")
	fmt.Scanln(&choice)

	switch strings.TrimSpace(choice) {
	case "1":
		fmt.Print("Enter folder path to zip: ")
		fmt.Scanln(&src)
		fmt.Print("Enter output zip name (e.g., output.zip): ")
		fmt.Scanln(&dst)
		if err := zipFolder(src, dst); err != nil {
			fmt.Println("❌ Error:", err)
		} else {
			fmt.Println("✅ Folder compressed to:", dst)
		}
	case "2":
		fmt.Print("Enter zip file path: ")
		fmt.Scanln(&src)
		fmt.Print("Enter destination folder: ")
		fmt.Scanln(&dst)
		if err := unzip(src, dst); err != nil {
			fmt.Println("❌ Error:", err)
		} else {
			fmt.Println("✅ Files extracted to:", dst)
		}
	default:
		fmt.Println("❌ Invalid option")
	}
}
