/**
 * Author: Amin Zamani
 * File: go-vtoken.go
 */

package vtoken

import (
	"time"

	"github.com/siaminz/go-vtoken/pkg/crypto"
	"github.com/siaminz/go-vtoken/pkg/token"
	"github.com/siaminz/go-vtoken/pkg/types"
)

var SECRET string = "ase&1*~#^2^#s0^=)^^7%b34"

// Returns a new token
func New(identifier string, expires time.Duration) (string, error) {
	data := types.New()
	data.SetIdentifier(identifier)
	data.SetExpires(expires)
	str := data.ToJSON()
	token, err := crypto.Encrypt(str, SECRET)
	if err != nil {
		return "", err
	}
	return token, nil
}

// Verifies a token
func Verify(token string) (*types.ID, error) {
	str, err := crypto.Decrypt(token, SECRET)
	if err != nil {
		return nil, err
	}
	id, err := types.FromJSON(str)
	if err != nil {
		return nil, err
	}
	return id, nil
}

// Generate a random string of A-Z, a-z, 0-9 characters with custom length, this tokens can not be verified.
func GenerateSimpleToken(length int) string {
	return token.Generate(length)
}

// Set Secret Key
func SetSecretKey(secret string) {
	// Secret Key must be 24 characters
	if len(secret) != 24 {
		panic("Secret Key must be 24 characters")
	}
	SECRET = secret
}
