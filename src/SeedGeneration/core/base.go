package core

import (
	"tmp/src/SeedGeneration/utils" // Import the utils package for utility functions
)

// SeedGenerator represents a seed generator.
type SeedGenerator struct {
	phrase string // The phrase used for seed generation
}

// NewSeedGenerator creates a new instance of SeedGenerator with the given phrase.
func NewSeedGenerator(phrase string) *SeedGenerator {
	return &SeedGenerator{phrase: phrase} // Initialize a new SeedGenerator instance with the provided phrase
}

// Generate generates a random seed.
func (sg *SeedGenerator) Generate() []byte {
	length := utils.GetRandomInt()                             // Get a random length for the seed
	randomBytes := make([]byte, length)                        // Generate random bytes of the specified length
	combinedBytes := append(randomBytes, []byte(sg.phrase)...) // Combine random bytes with the phrase
	return utils.HashDigest(combinedBytes)                     // Hash the combined bytes and return the result
}
