package core

import (
	"crypto/sha256"           // Import package to use SHA-256 hash function
	"encoding/hex"            // Import package for hexadecimal encoding
	"tmp/src/HMAC/algorithms" // Import package for algorithm definitions
)

// HMACClient represents HMAC (Hash-based Message Authentication Code) client for message encryption and decryption.
type HMACClient struct {
	Algorithm   string            // Algorithm specifies the HMAC algorithm being used
	Secret      []byte            // Secret stores the secret key used for HMAC
	SymbolCount int               // SymbolCount specifies the number of symbols per chunk
	DecryptDict map[string]string // DecryptDict stores precomputed HMAC hashes for decryption
}

// NewHMACClient creates a new instance of HMACClient with the given parameters.
func NewHMACClient(algorithm string, secret []byte, symbolCount int) *HMACClient {
	return &HMACClient{
		Algorithm:   algorithm,
		Secret:      secret,
		SymbolCount: symbolCount,
		DecryptDict: make(map[string]string), // Initialize decryption dictionary
	}
}

// InitDecryptDict initializes the decryption dictionary with all possible combinations of characters.
func (h *HMACClient) InitDecryptDict() {
	asciiLetters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// digits
	digits := "0123456789"

	// punctuation
	punctuation := `!"#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + "{|}~"

	// Concatenate them together
	allChars := asciiLetters + digits + punctuation
	combinations := h.GenerateCombinations(allChars, h.SymbolCount)

	for _, comb := range combinations {
		key := h.EncryptMessage(comb)
		h.DecryptDict[key] = comb
	}
}

// GenerateCombinations generates all combinations of characters with specified symbol count.
func (h *HMACClient) GenerateCombinations(s string, k int) []string {
	var combinations []string
	if k == 0 {
		return []string{""}
	}
	for i := 0; i < len(s); i++ {
		for _, comb := range h.GenerateCombinations(s, k-1) {
			combinations = append(combinations, string(s[i])+comb)
		}
	}
	return combinations
}

// EncryptMessageByChunks encrypts a message by dividing it into chunks and encrypting each chunk.
func (h *HMACClient) EncryptMessageByChunks(message string) string {
	var encryptedMessage string
	for i := 0; i < len(message); i += h.SymbolCount {
		end := i + h.SymbolCount
		if end > len(message) {
			end = len(message)
		}
		chunk := message[i:end]
		encryptedMessage += h.EncryptMessage(chunk)
	}
	return encryptedMessage
}

// EncryptMessage encrypts a message using HMAC.
func (h *HMACClient) EncryptMessage(message string) string {
	hash := HMAC.New(sha256.New, h.Secret)   // Create new HMAC hash using SHA-256 and secret key
	hash.Write([]byte(message))              // Write message to hash
	return hex.EncodeToString(hash.Sum(nil)) // Return hexadecimal encoded hash
}

// DecryptMessageByChunks decrypts a message by dividing it into chunks and decrypting each chunk.
func (h *HMACClient) DecryptMessageByChunks(message string) string {
	var msgRaw string
	var chank_size = algorithms.HashTypeAndLength[h.Algorithm] // Get chunk size based on algorithm
	if len(message)%h.SymbolCount == 0 {
		for i := 0; i < len(message); i += chank_size {
			end := i + chank_size
			if end > len(message) {
				end = len(message)
			}
			chunk := message[i:end]
			msgRaw += h.DecryptMessage(chunk)
		}
		return msgRaw
	}
	panic("The algorithm is invalid")
}

// DecryptMessage decrypts a message using HMAC.
func (h *HMACClient) DecryptMessage(message string) string {
	if val, ok := h.DecryptDict[message]; ok { // Check if HMAC hash exists in decryption dictionary
		return val
	}
	panic("The algorithm is invalid")
}
