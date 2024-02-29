package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func GenerateKeyPair(password string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	// Generate a private key using crypto/rand.Reader for randomness
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}

	// Compute the SHA-256 hash of the password
	hash := sha256.Sum256([]byte(password))

	// Derive a new private key using the hashed password as an additional seed
	privateKey.D = new(big.Int).SetBytes(hash[:])

	// Compute the corresponding public key
	publicKey := &privateKey.PublicKey

	return privateKey, publicKey, nil
}
