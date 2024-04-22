package utils

import (
	"encoding/base64" // Import package for base64 encoding and decoding
	"fmt"             // Import package for formatted I/O
	"math/big"        // Import package for big integer arithmetic
)

// bytesToInt converts any value to an integer from the big endian bytes representation
func bytesToInt(value interface{}) *big.Int {
	bytes := toBytes(value)              // Convert value to bytes
	return big.NewInt(0).SetBytes(bytes) // Create a new big integer from bytes
}

// IntToBytes converts an integer value to bytes in big endian representation
func IntToBytes(value *big.Int) []byte {
	byteLen := (value.BitLen() + 7) / 8 // Calculate byte length based on bit length
	bytes := make([]byte, byteLen)      // Create byte slice of appropriate length
	return value.FillBytes(bytes)       // Fill byte slice with the integer's bytes in big endian representation
}

// dataToB64Str encodes data in base64, optionally stripping padding
func dataToB64Str(data interface{}, strip bool) string {
	bytes := toBytes(data)                          // Convert data to bytes
	b64 := base64.StdEncoding.EncodeToString(bytes) // Encode bytes to base64 string
	if strip {
		b64 = stripPadding(b64) // Strip padding if requested
	}
	return b64 // Return base64 encoded string
}

// b64d decodes base64 to bytes, appends padding just in case
func b64d(data string, pad bool) ([]byte, error) {
	bytes, err := base64.StdEncoding.DecodeString(data) // Decode base64 string to bytes
	if err != nil {
		return nil, err
	}
	if pad {
		bytes = append(bytes, []byte("===")...) // Append padding if requested
	}
	return bytes, nil // Return decoded bytes
}

// toBytes converts data to bytes representation
func toBytes(data interface{}) []byte {
	switch v := data.(type) {
	case string:
		return []byte(v) // Convert string to bytes
	case []byte:
		return v // Return bytes if already in byte slice
	case int:
		return IntToBytes(big.NewInt(int64(v))) // Convert integer to bytes
	default:
		fmt.Println("UNTYPED:", data)          // Print message for unhandled types
		return []byte(fmt.Sprintf("%v", data)) // Convert other types to string and then to bytes
	}
}

// stripPadding strips padding from base64 encoded string
func stripPadding(b64 string) string {
	for i := len(b64) - 1; i >= 0; i-- {
		if b64[i] == '=' {
			continue
		}
		return b64[:i+1] // Return string without padding
	}
	return "" // Return empty string if no padding found
}
