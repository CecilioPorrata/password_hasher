package utils

import "fmt"

func CombineHashes(RandIndex, RandLength int, privKey, pubKey string) string {
	// Here you want to take in the RandomIndex number, and use that to start an insert
	// of the pubKey into the privKey by 'overwriting'

	// Check if the length of the insertion is larger than the length available
	// i.e. if we have a randIndex of 63, and we have an insertion amount(randLength)
	// of 10, then we don't have enough room.

	// Error check for edge case (This should not happen)
	if RandIndex < 0 {
		fmt.Println("Error: index is too low")
	}
	// To make it simple here I will just -10 until it's good
	if RandIndex > 64-RandLength {
		RandIndex = (64 - RandLength) - 10
	}

	// We take a slice of the private key based on RandIndes and RandLength
	PrivKeySliceToInsert := privKey[RandIndex : RandIndex+RandLength]
	// We simply concatenate them together to form the NewPubKey
	NewPubKey := pubKey[:RandIndex] + PrivKeySliceToInsert + pubKey[RandIndex+RandLength:]

	fmt.Println("Random Index: ", RandIndex)
	fmt.Println("Random Length: ", RandLength)
	fmt.Println("New Public Key", NewPubKey)
	return NewPubKey
}
