package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// generate a random number between 1 and 100. the +1 shifts the int by 1
	min, max := 1, 100

	n := rand.IntN(max-min+1) + min

	fmt.Println("The random int is: ", n)

}
