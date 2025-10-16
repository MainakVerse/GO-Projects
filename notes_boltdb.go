package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("notes.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	reader := bufio.NewReader(os.Stdin)
	bucket := []byte("Notes")

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket)
		return err
	})

	fmt.Println("🗒️  Persistent Notes App (BoltDB)")
	fmt.Println("----------------------------------")

	for {
		fmt.Print("\nCommand (add/list/del/exit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(strings.ToLower(cmd))

		switch cmd {
		case "add":
			fmt.Print("Note title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Note content: ")
			content, _ := reader.ReadString('\n')
			content = strings.TrimSpace(content)

			db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket(bucket)
				return b.Put([]byte(title), []byte(content))
			})
			fmt.Println("✅ Note saved.")

		case "list":
			db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket(bucket)
				fmt.Println("\n📋 Your Notes:")
				b.ForEach(func(k, v []byte) error {
					fmt.Printf("- %s: %s\n", k, v)
					return nil
				})
				return nil
			})

		case "del":
			fmt.Print("Note title to delete: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket(bucket)
				return b.Delete([]byte(title))
			})
			fmt.Println("🗑️ Note deleted (if existed).")

		case "exit":
			fmt.Println("👋 Goodbye!")
			return

		default:
			fmt.Println("❌ Invalid command.")
		}
	}
}
