package main

import (
	"fmt"
)

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}

	fibSeries := make([]int, n)
	fibSeries[0] = 0
	if n > 1 {
		fibSeries[1] = 1
	}

	for i := 2; i < n; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	return fibSeries
}

func main() {
	var n int
	fmt.Print("Enter the number of Fibonacci numbers to generate: ")
	fmt.Scan(&n)

	fibNumbers := fibonacci(n)
	fmt.Println("Fibonacci Series:", fibNumbers)
}
