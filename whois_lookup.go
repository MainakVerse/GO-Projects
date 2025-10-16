package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func whoisLookup(domain string) (string, error) {
	conn, err := net.DialTimeout("tcp", "whois.verisign-grs.com:43", 5*time.Second)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return "", err
	}

	response := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		response += scanner.Text() + "\n"
	}
	return response, scanner.Err()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter domain or IP: ")
	domain, _ := reader.ReadString('\n')
	domain = strings.TrimSpace(domain)

	if domain == "" {
		fmt.Println("❌ Please enter a valid domain.")
		return
	}

	fmt.Println("\n🔍 Performing WHOIS lookup...\n")
	result, err := whoisLookup(domain)
	if err != nil {
		fmt.Println("❌ Error:", err)
		return
	}

	fmt.Println(result)
}
