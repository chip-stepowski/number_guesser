package main

import (
	"fmt"
	"math/rand/v2"
)

// declare the guess function with return type of int
func guess() int {
	var i int

	fmt.Println("Enter a number between 1 and 100:")

	fmt.Scan(&i)

	return i
}

func main() {

	// generate a random number between 1 and 100. the +1 shifts the int by 1
	min, max := 1, 100

	n := rand.IntN(max-min+1) + min

	fmt.Println("The random int is: ", n)

	numOfGuesses := 0

	for {
		// player guesses
		playerGuess := guess()

		fmt.Println("Player guess is:", playerGuess)

		// program should reply higher or lower
		// if less than
		if playerGuess < n {
			fmt.Println("Too low, guess higher!")
			numOfGuesses++
		}
		// if greater than
		if playerGuess > n {
			fmt.Println("Too high, guess lower!")
			numOfGuesses++
		}
		// if equal
		if playerGuess == n {
			fmt.Println("You win!")
			numOfGuesses++
			break
		}
	}
	// count # of guesses and show total at the end
	fmt.Println("Number of guesses:", numOfGuesses)
}
