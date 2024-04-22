package utils

import (
	"crypto/elliptic" // Package for elliptic curve cryptography
	"crypto/rand"     // Package for secure random number generation
	"crypto/sha256"   // Package for SHA-256 hashing algorithm
	"fmt"             // Package for formatted I/O
	"math/big"        // Package for arbitrary-precision arithmetic
)

// hashTypes defines supported hash algorithms
var hashTypes = map[string]func([]byte) []byte{
	"sha3_256": func(data []byte) []byte { // Define SHA-256 hashing function
		hash := sha256.Sum256(data) // Compute SHA-256 hash
		return hash[:]              // Return hash as byte slice
	},
}

// GenerateSalt generates a random salt of the specified size in bytes
func GenerateSalt(size int) []byte {
	salt := make([]byte, size) // Create byte slice for salt
	_, err := rand.Read(salt)  // Fill byte slice with random data
	if err != nil {
		// Handle error if occurred during salt generation
		fmt.Println("Error generating salt:", err)
	}
	return salt // Return generated salt
}

// CurveByName gets a curve by name, case-insensitive
func CurveByName(name string) elliptic.Curve {
	switch name {
	case "secp256k1":
		return elliptic.P256() // Placeholder, replace with actual curve implementation
	case "curve25519":
		return nil // Placeholder, replace with actual curve implementation
	default:
		return nil
	}
}

// mod returns a mod b, accounting for positive/negative numbers
func mod(a, b *big.Int) *big.Int {
	mod := new(big.Int).Mod(a, b) // Calculate a mod b
	if mod.Sign() < 0 {           // If result is negative
		mod.Add(mod, b) // Add b to make it positive
	}
	return mod // Return positive modulus
}

// hashData converts all provided values to bytes, and returns the digest in bytes
func hashData(values ...interface{}) []byte {
	var concatenated []byte
	for _, value := range values { // Iterate over values
		bytes := toBytes(value)                       // Convert value to bytes
		concatenated = append(concatenated, bytes...) // Concatenate bytes
	}
	return hashTypes["sha3_256"](concatenated) // Return hash of concatenated bytes
}

// hashNumeric computes the cryptographic hash of the provided values and returns the digest in integer form
func hashNumeric(values ...interface{}) *big.Int {
	hash := hashData(values...) // Compute hash of values
	return bytesToInt(hash)     // Convert hash bytes to big integer
}
