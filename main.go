package main

import (
	"fmt"

	"pwhasher/utils"
)

func main() {
	password := "mysecretpassword"
	utils.GenerateKeyPair(password)

	privateKey, publicKey, err := utils.GenerateKeyPair(password)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	privKey := publicKey.Y
	pubKey := publicKey.X
	// Keep on file or discard
	fmt.Println("Private Key:", privateKey.D)
	// Send to User
	fmt.Println("Public Key:", pubKey) // publicKey.X
	// Send to Database
	fmt.Println("Private Key:", privKey) // publicKey.Y

	utils.CombineHashes(utils.RandomizerIndex(), utils.RandomizerLength(), privKey.String(), pubKey.String())

}
