package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

// declare the guess function with return type of int
func guess() int {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a number between 1 and 100:")
		if !scanner.Scan() {
			return 0
		}

		input := strings.TrimSpace(scanner.Text())
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}
		if num > 100 || num < 1 {
			fmt.Println("Invalid input, input is smaller than 1 or greater than 100.")
			continue
		}
		return num
	}
}

func main() {

	// generate a random number between 1 and 100. the +1 shifts the int by 1
	min, max := 1, 100

	n := rand.IntN(max-min+1) + min

	numOfGuesses := 0

	for {
		numOfGuesses++
		// player guesses
		playerGuess := guess()

		fmt.Println("Player guess is:", playerGuess)

		// program should reply higher or lower
		// if less than
		if playerGuess < n {
			fmt.Println("Too low, guess higher!")
		}
		// if greater than
		if playerGuess > n {
			fmt.Println("Too high, guess lower!")
		}
		// if equal
		if playerGuess == n {
			fmt.Println("You win!")
			break
		}
	}
	// count # of guesses and show total at the end
	fmt.Println("Number of guesses:", numOfGuesses)
}
