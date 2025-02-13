package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int

	// Accept user input
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)
	fmt.Print("Enter third number: ")
	fmt.Scan(&num3)

	// Determine the largest number
	largest := num1

	if num2 > largest {
		largest = num2
	}
	if num3 > largest {
		largest = num3
	}

	fmt.Printf("The largest number is: %d\n", largest)
}
