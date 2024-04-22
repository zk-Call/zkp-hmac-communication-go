package utils

import (
	"encoding/json" // Package for JSON encoding and decoding
)

// dumpObject dumps a JSON Dataclass to compressed JSON
func dumpObject(dc interface{}) ([]byte, error) {
	return json.Marshal(dc) // Marshal the JSON Dataclass into compressed JSON bytes
}
