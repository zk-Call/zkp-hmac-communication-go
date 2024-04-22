package utils

// ToBytes converts data to bytes based on its type.
func ToBytes(data interface{}, encoding string) []byte {
	switch v := data.(type) {
	case []byte: // If data is already a byte slice, return it directly
		return v
	case string: // If data is a string, convert it to bytes and return
		return []byte(v)
	default:
		return nil // Return nil if data is neither a byte slice nor a string
	}
}

// ToStr converts bytes to a string.
func ToStr(data []byte, encoding string) string {
	return string(data) // Convert bytes to string and return
}
