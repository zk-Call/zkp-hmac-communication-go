package models

import (
	"crypto/elliptic" // Import elliptic curve functions
	"encoding/json"   // Import package for JSON encoding and decoding
	"math/big"        // Import package for big integer arithmetic
)

// Define ZeroKnowledgeParams struct
type ZeroKnowledgeParams struct {
	Algorithm string // Algorithm used for zero-knowledge proofs
	Curve     string // Elliptic curve used for zero-knowledge proofs
	Salt      []byte // Salt value for zero-knowledge proofs
}

// Define ZeroKnowledgeSignature struct
type ZeroKnowledgeSignature struct {
	Params    ZeroKnowledgeParams // Parameters for zero-knowledge proofs
	Signature []byte              // Signature data
}

// Define ZeroKnowledgeProof struct
type ZeroKnowledgeProof struct {
	Params ZeroKnowledgeParams // Parameters for zero-knowledge proofs
	C      []byte              // Proof data
	M      []byte              // Proof data
}

// Define ZeroKnowledgeData struct
type ZeroKnowledgeData struct {
	Data  string             // Data for zero-knowledge proofs
	Proof ZeroKnowledgeProof // Proof associated with the data
}

// Define Curve struct
type Curve struct {
	elliptic.Curve // Elliptic curve object
}

// Define Point struct
type Point struct {
	X, Y *big.Int // X and Y coordinates of a point
}

// ToJSON converts ZeroKnowledgeParams to JSON
func (params *ZeroKnowledgeParams) ToJSON() ([]byte, error) {
	return json.Marshal(params) // Convert struct to JSON bytes
}

// FromJSON converts JSON data to ZeroKnowledgeParams
func (params *ZeroKnowledgeParams) FromJSON(data []byte) error {
	return json.Unmarshal(data, params) // Parse JSON bytes into struct
}

// ToJSON converts ZeroKnowledgeSignature to JSON
func (signature *ZeroKnowledgeSignature) ToJSON() ([]byte, error) {
	return json.Marshal(signature) // Convert struct to JSON bytes
}

// FromJSON converts JSON data to ZeroKnowledgeSignature
func (signature *ZeroKnowledgeSignature) FromJSON(data []byte) error {
	return json.Unmarshal(data, signature) // Parse JSON bytes into struct
}

// ToJSON converts ZeroKnowledgeProof to JSON
func (proof *ZeroKnowledgeProof) ToJSON() ([]byte, error) {
	return json.Marshal(proof) // Convert struct to JSON bytes
}

// FromJSON converts JSON data to ZeroKnowledgeProof
func (proof *ZeroKnowledgeProof) FromJSON(data []byte) error {
	return json.Unmarshal(data, proof) // Parse JSON bytes into struct
}

// ToJSON converts ZeroKnowledgeData to JSON
func (data *ZeroKnowledgeData) ToJSON() ([]byte, error) {
	return json.Marshal(data) // Convert struct to JSON bytes
}

// FromJSON converts JSON data to ZeroKnowledgeData
func (data *ZeroKnowledgeData) FromJSON(dataBytes []byte) error {
	return json.Unmarshal(dataBytes, data) // Parse JSON bytes into struct
}
