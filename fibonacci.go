package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	seq := []int{0, 1}
	for i := 2; i < n; i++ {
		next := seq[i-1] + seq[i-2]
		seq = append(seq, next)
	}
	return seq
}

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("❌ Enter a positive number.")
		return
	}

	fmt.Printf("\n🌀 First %d Fibonacci Numbers:\n", n)
	for _, val := range fibonacci(n) {
		fmt.Printf("%d ", val)
	}
	fmt.Println("\n✅ Done!")
}
