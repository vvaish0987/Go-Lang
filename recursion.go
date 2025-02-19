package main

import "fmt"

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
	} else {
		fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
	}
}
