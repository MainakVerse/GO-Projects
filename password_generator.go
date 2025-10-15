package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func generatePassword(length int, charset string) string {
	password := make([]byte, length)
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	return string(password)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	charsets := map[string]string{
		"1": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"2": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
		"3": "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+",
	}

	for {
		fmt.Print("\nEnter password length (e.g. 12): ")
		lengthStr, _ := reader.ReadString('\n')
		lengthStr = strings.TrimSpace(lengthStr)
		length, err := strconv.Atoi(lengthStr)
		if err != nil || length <= 0 {
			fmt.Println("❌ Invalid length")
			continue
		}

		fmt.Println("Choose strength: 1=Letters | 2=Letters+Digits | 3=Strong (All)")
		fmt.Print("Enter choice: ")
		strength, _ := reader.ReadString('\n')
		strength = strings.TrimSpace(strength)

		charset, ok := charsets[strength]
		if !ok {
			fmt.Println("❌ Invalid choice")
			continue
		}

		fmt.Println("🔑 Generated Password:", generatePassword(length, charset))

		fmt.Print("Generate another? (y/n): ")
		ans, _ := reader.ReadString('\n')
		if strings.TrimSpace(strings.ToLower(ans)) != "y" {
			fmt.Println("👋 Goodbye!")
			break
		}
	}
}
