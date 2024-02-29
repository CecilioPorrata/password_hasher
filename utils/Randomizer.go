package utils

import (
	"math/rand"
)

// Returns random numbers for use throughout the program

// Create a random number between 0 and 10
// This denotes the amount of characters that will be used to inserted
func RandomizerLength() int {
	// Generates a new random int
	RandLength := rand.Intn(11)
	return int(RandLength)
}

// Generates Random number from 0-64
// Denotes the index location where insertion will start
func RandomizerIndex() int {

	RandIndex := rand.Intn(65)
	return RandIndex
}
