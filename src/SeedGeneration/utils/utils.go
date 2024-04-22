package utils

import (
	"crypto/sha256" // Import the SHA-256 cryptographic hash function
	"math/rand"     // Import the math/rand package for random number generation
)

// GetRandomInt generates a random integer.
func GetRandomInt() int {
	randomInt := rand.Intn(1 << 20) // Generate a random integer in the range [0, 1 << 20)
	return randomInt
}

// HashDigest calculates the SHA-256 hash of the given byte slice.
func HashDigest(combinedBytes []byte) []byte {
	hash := sha256.Sum256(combinedBytes) // Calculate the SHA-256 hash of the byte slice
	return hash[:]                       // Return the hash as a byte slice
}
