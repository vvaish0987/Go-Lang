package main

import (
	"fmt"
)

func decimalToBinary(n int) string {
	if n == 0 {
		return "0" // Special case: 0 in binary is 0
	}

	binary := "" // Store the binary representation as a string

	// Loop until the decimal number is greater than 0
	for n > 0 {
		remainder := n % 2 // Get the remainder when divided by 2
		binary = fmt.Sprintf("%d%s", remainder, binary) // Prepend the remainder to the binary string
		n = n / 2 // Divide the number by 2
	}

	return binary
}

func main() {
	var decimal int

	// Prompt the user to enter a decimal number
	fmt.Print("Enter a decimal number: ")
	fmt.Scan(&decimal)

	// Convert the decimal number to binary
	binary := decimalToBinary(decimal)

	// Print the result
	fmt.Printf("Binary representation of %d is %s\n", decimal, binary)
}