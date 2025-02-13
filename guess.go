package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int
	attempts := 0

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I have selected a number between 1 and 100. Try to guess it!")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter your guess: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		var err error
		guess, err = strconv.Atoi(input)
		if err != nil || guess < 1 || guess > 100 {
			fmt.Println("Invalid input! Please enter a number between 1 and 100.")
			continue
		}

		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d in %d attempts.\n", target, attempts)
			break
		}
	}
}
