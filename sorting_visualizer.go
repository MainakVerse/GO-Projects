package main

import (
	"fmt"
	"time"
)

func printArray(arr []int, highlight int) {
	for i, v := range arr {
		if i == highlight {
			fmt.Printf("[%d] ", v)
		} else {
			fmt.Printf("%d ", v)
		}
	}
	fmt.Println()
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			printArray(arr, j)
			time.Sleep(200 * time.Millisecond)
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println("✅ Sorted:", arr)
}

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			printArray(arr, j)
			time.Sleep(200 * time.Millisecond)
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	fmt.Println("✅ Sorted:", arr)
}

func main() {
	arr := []int{5, 3, 8, 1, 4}
	fmt.Println("📊 Sorting Visualizer CLI")
	fmt.Println("---------------------------")
	fmt.Println("Array:", arr)

	var choice int
	fmt.Print("\nChoose algorithm (1=Bubble, 2=Selection): ")
	fmt.Scan(&choice)

	if choice == 1 {
		bubbleSort(arr)
	} else {
		selectionSort(arr)
	}
}
