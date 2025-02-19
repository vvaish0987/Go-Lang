package main

import "fmt"

func main() {
	var n, sum int

	fmt.Print("Enter a number N: ")
	fmt.Scan(&n)

	for i := 2; i <= n; i += 2 {
		sum += i
	}

	fmt.Println("Sum of even numbers from 1 to", n, "is:", sum)
}
