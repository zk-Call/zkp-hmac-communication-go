package algorithms

import (
	"crypto/md5"               // Import package for MD5 hash function
	"crypto/sha1"              // Import package for SHA-1 hash function
	"crypto/sha256"            // Import package for SHA-256 hash function
	"crypto/sha512"            // Import package for SHA-512 hash function
	"golang.org/x/crypto/sha3" // Import package for SHA-3 hash function
	"hash"                     // Import package for hash functions
)

// HashTypes defines supported hash algorithms along with their corresponding hash functions.
var HashTypes = map[string]func() hash.Hash{
	"md5":      md5.New,       // MD5 hash function
	"sha1":     sha1.New,      // SHA-1 hash function
	"sha224":   sha256.New224, // SHA-224 hash function
	"sha256":   sha256.New,    // SHA-256 hash function
	"sha512":   sha512.New,    // SHA-512 hash function
	"sha3_224": sha3.New224,   // SHA3-224 hash function
	"sha3_256": sha3.New256,   // SHA3-256 hash function
	"sha3_384": sha3.New384,   // SHA3-384 hash function
	"sha3_512": sha3.New512,   // SHA3-512 hash function
}

// JWTAlgorithms defines JWT (JSON Web Token) algorithms along with their corresponding hash functions.
var JWTAlgorithms = map[string]func() hash.Hash{
	"HS3_224": sha3.New224, // SHA3-224 hash function for JWT
	"HS3_256": sha3.New256, // SHA3-256 hash function for JWT
	"HS3_384": sha3.New384, // SHA3-384 hash function for JWT
	"HS3_512": sha3.New512, // SHA3-512 hash function for JWT
}
