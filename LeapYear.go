package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	// A year is a leap year if it is divisible by 4,
	// but not by 100, unless it is also divisible by 400.
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}

func main() {
	var year int

	// Prompt the user to enter a year
	fmt.Print("Enter a year: ")
	fmt.Scan(&year)

	// Check if the year is a leap year
	if isLeapYear(year) {
		fmt.Printf("%d is a leap year.\n", year)
	} else {
		fmt.Printf("%d is not a leap year.\n", year)
	}
}