package core

import (
	"crypto/ecdsa"                            // Import ECDSA cryptographic functions
	"crypto/elliptic"                         // Import elliptic curve functions
	"crypto/rand"                             // Import cryptographic random number generator
	"crypto/sha256"                           // Import SHA-256 cryptographic hash function
	"encoding/json"                           // Import package for JSON encoding and decoding
	"errors"                                  // Import package for error handling
	"github.com/golang-jwt/jwt/v4"            // Import JWT package for JSON Web Tokens
	"math/big"                                // Import package for big integer arithmetic
	"time"                                    // Import package for handling time
	zkx_models "tmp/src/ZeroKnowledge/models" // Import Zero Knowledge models
	zkx_utils "tmp/src/ZeroKnowledge/utils"   // Import Zero Knowledge utility functions
)

// Define ZeroKnowledge struct
type ZeroKnowledge struct {
	Params    zkx_models.ZeroKnowledgeParams // Parameters for Zero Knowledge
	Curve     zkx_models.Curve               // Elliptic curve
	Bits      int                            // Number of bits
	Secret    []byte                         // Secret key for JWT
	Algorithm string                         // JWT algorithm
	Issuer    string                         // Issuer for JWT
}

// New creates a new instance of ZeroKnowledge
func New(curveName string, hashAlg string, jwtSecret []byte, jwtAlg string, saltSize int) (*ZeroKnowledge, error) {
	// Get the elliptic curve object
	curve := zkx_utils.CurveByName(curveName)
	if curve == nil {
		return nil, errors.New("Invalid Curve Name")
	}

	// Initialize ZeroKnowledgeParams object
	params := zkx_models.ZeroKnowledgeParams{
		Algorithm: hashAlg,
		Curve:     curveName,
		Salt:      zkx_utils.GenerateSalt(saltSize),
	}

	// Create a new instance of ZeroKnowledge
	zk := ZeroKnowledge{
		Params:    params,
		Curve:     curve,
		Secret:    jwtSecret,
		Algorithm: jwtAlg,
	}

	return &zk
}

// GenerateJWT generates a JSON Web Token (JWT) using the provided signature and expiration time
func (z *ZeroKnowledge) GenerateJWT(signature zkx_models.ZeroKnowledgeSignature, exp time.Duration) (string, error) {
	if len(z.Secret) == 0 {
		return "", errors.New("JWT secret is empty")
	}
	now := time.Now().UTC()
	claims := map[string]interface{}{
		"signature": signature,
		"iat":       now,
		"nbf":       now,
		"exp":       now.Add(exp),
		"iss":       z.Issuer,
	}
	token, err := JwtEncode(claims, z.Secret, z.Algorithm)
	if err != nil {
		return "", err
	}
	return token, nil
}

// JwtEncode encodes JWT claims using the provided secret and algorithm
func JwtEncode(claims map[string]interface{}, secret []byte, algorithm string) (string, error) {
	token, err := jwt.Encode(claims, secret, algorithm)
	if err != nil {
		return "", err
	}
	return string(token), nil
}

// verifyJWT verifies a JSON Web Token (JWT) and returns decoded data if valid
func (z *ZeroKnowledge) verifyJWT(tok []byte) (map[string]interface{}, error) {
	if len(z.Secret) == 0 {
		return nil, errors.New("JWT secret is empty")
	}
	return JwtDecode(tok, z.Secret, z.Issuer, z.Algorithm)
}

// JwtDecode decodes a JWT using the provided secret, issuer, and algorithm
func JwtDecode(tok []byte, secret []byte, issuer string, algorithm string) (map[string]interface{}, error) {
	claims, err := jwt.Parse(string(tok), secret, issuer, algorithm)
	if err != nil {
		return nil, err
	}
	return claims, nil
}

// newPoint creates a new Point object
func (z *ZeroKnowledge) NewPoint(value interface{}) zkx_models.Point {
	switch v := value.(type) {
	case int:
		// Convert the integer value to big.Int
		bigIntValue := new(big.Int).SetInt64(int64(v))
		// Multiply the generator point of the curve by the integer value
		x, y := z.Curve.ScalarBaseMult(bigIntValue.Bytes())
		return zkx_models.Point{X: x, Y: y}
	case []byte:
		// Convert the byte slice to big.Int
		bigIntValue := new(big.Int).SetBytes(v)
		// Multiply the generator point of the curve by the byte slice
		x, y := z.Curve.ScalarBaseMult(bigIntValue.Bytes())
		return zkx_models.Point{X: x, Y: y}
	case zkx_models.ZeroKnowledgeSignature:
		// Assuming the value is already a point
		return zkx_models.Point{X: value.X, Y: value.Y}
	default:
		// Handle other types if necessary
		return zkx_models.Point{} // Return an empty point as default
	}
}

// createSignature creates a signature object using the provided secret key
func (z *ZeroKnowledge) CreateSignature(secret []byte) zkx_models.ZeroKnowledgeSignature {
	signature := z.Hash(secret)
	return zkx_models.ZeroKnowledgeSignature{
		Params:    z.Params,
		Signature: z.NewPoint(signature * z.Curve.Params().Gx),
	}
}

// createProof creates a proof object using the provided secret key and optional data
func (z *ZeroKnowledge) CreateProof(secret []byte, data interface{}) zkx_models.ZeroKnowledgeProof {
	key := z.Hash(secret)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(z.Bits)))
	Rx, Ry := z.Curve.ScalarBaseMult(r.Bytes())
	R := zkx_models.Point{Rx, Ry}
	c := z.Hash(data, R)
	m := new(big.Int).Mod(new(big.Int).Sub(r, new(big.Int).Mul(c, key)), z.Curve.Params().N)
	return zkx_models.ZeroKnowledgeProof{
		Params: z.Params,
		C:      zkx_utils.IntToBytes(c),
		M:      zkx_utils.IntToBytes(m),
	}
}

// hash hashes the values provided modulo the curve order
func (z *ZeroKnowledge) Hash(values ...interface{}) *big.Int {
	// Concatenate all values into a single byte slice
	var concatenated []byte
	for _, value := range values {
		switch v := value.(type) {
		case int:
			concatenated = append(concatenated, zkx_utils.IntToBytes(v)...)
		case string:
			concatenated = append(concatenated, []byte(v)...)
		case []byte:
			concatenated = append(concatenated, v...)
		default:
			panic(errors.New("Unknown type"))
		}
	}

	// Calculate the hash of the concatenated byte slice
	hash := sha256.Sum256(concatenated)

	// Convert the hash to a big.Int
	hashInt := new(big.Int).SetBytes(hash[:])

	// Reduce the hash modulo the curve order
	return hashInt.Mod(hashInt, z.Curve.Params().N)
}

// _toPoint converts a value to a point on the elliptic curve

func (z *ZeroKnowledge) _toPoint(value interface{}) zkx_models.Point {
	// Convert the value to a big integer
	var scalar *big.Int
	switch v := value.(type) {
	case int:
		scalar = big.NewInt(int64(v))
	case *big.Int:
		scalar = v
	default:
		// Handle other types if necessary
		panic(errors.New("Unknown type"))
	}

	// Get the base point of the curve

	// Multiply the base point by the scalar to get the resulting point
	x, y := z.Curve.ScalarBaseMult(scalar.Bytes())

	// Create and return the zkx_models.Point object
	return zkx_models.Point{X: x, Y: y}

}

// Verify verifies a challenge against a signature and optional data
func (z *ZeroKnowledge) Verify(challenge interface{}, signature zkx_models.ZeroKnowledgeSignature, data interface{}) bool {
	// Convert the challenge to the appropriate type
	var challengeData []byte
	switch c := challenge.(type) {
	case zkx_models.ZeroKnowledgeData:
		challengeData = []byte(c.Data)
	case zkx_models.ZeroKnowledgeProof:
		// Handle if necessary
	default:
		// Handle other cases if needed
	}

	// Compute the hash of the challenge data
	hash := z.Hash(challengeData)

	// Convert the signature to ecdsa.PublicKey
	x, y := elliptic.Unmarshal(z.Curve.Params(), signature.Signature)
	publicKey := ecdsa.PublicKey{Curve: z.Curve, X: x, Y: y}

	// Verify the signature using the ecdsa.Verify function
	return ecdsa.Verify(&publicKey, hash, signature.Signature)
}

// Sign creates a ZeroKnowledgeData object with a proof for the provided data
func (z *ZeroKnowledge) Sign(secret []byte, data interface{}) *zkx_models.ZeroKnowledgeData {
	proof := z.CreateProof(secret, data) // Create proof for the data

	return &zkx_models.ZeroKnowledgeData{
		Data:  data,
		Proof: proof,
	}
}

// Token generates a random token of specified length in bytes
func Token(z ZeroKnowledge) ([]byte, error) {
	bytes := (z.Bits + 7) >> 3 // Calculate number of bytes based on bits
	token := make([]byte, bytes)
	_, err := rand.Read(token) // Generate random bytes
	if err != nil {
		return nil, err
	}
	return token, nil
}

// Login performs a login using the provided login data
func (z *ZeroKnowledge) Login(loginData zkx_models.ZeroKnowledgeData) bool {
	data, err := z.verifyJWT([]byte(loginData.Data))
	if err != nil || data == nil {
		return false
	}
	signature := zkx_models.ZeroKnowledgeSignature{}
	if err := json.Unmarshal([]byte(data["signature"].(string)), &signature); err != nil {
		return false
	}
	return z.Verify(loginData, signature, nil)
}
