package models

import "time"

// https://mholt.github.io/json-to-go/
type SigningKey struct {
	PrivateKey string    `json:"private_key"`
	PublicKey  string    `json:"public_key"`
	NotBefore  time.Time `json:"not_before"`
	NotAfter   time.Time `json:"not_after"`
	Password   string    `json:"password"`
	Kid        string    `json:"kid"`
	PublicJwk  struct {
		Alg string `json:"alg"`
		Crv string `json:"crv"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Use string `json:"use"`
		X   string `json:"x"`
		Y   string `json:"y"`
	} `json:"public_jwk"`
	PrivateJwk struct {
		Alg string `json:"alg"`
		Crv string `json:"crv"`
		D   string `json:"d"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Use string `json:"use"`
		X   string `json:"x"`
		Y   string `json:"y"`
	} `json:"private_jwk"`
}
