package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return // Closed or filtered
	}
	conn.Close()
	fmt.Printf("✅ Port %d is open\n", port)
}

func main() {
	var host string
	fmt.Print("Enter host (e.g. scanme.nmap.org or 127.0.0.1): ")
	fmt.Scanln(&host)

	start, end := 1, 1024 // default range
	fmt.Printf("\n🔍 Scanning %s (ports %d–%d)...\n", host, start, end)

	var wg sync.WaitGroup
	for port := start; port <= end; port++ {
		wg.Add(1)
		go scanPort(host, port, &wg)
	}

	wg.Wait()
	fmt.Println("\n✅ Scan complete!")
}
