package main

import (
	"fmt"
)

// check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var limit int
	fmt.Print("Generate primes up to: ")
	fmt.Scan(&limit)

	if limit < 2 {
		fmt.Println("❌ Enter a number ≥ 2")
		return
	}

	fmt.Printf("\n🔢 Prime Numbers up to %d:\n", limit)
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\n✅ Done!")
}
