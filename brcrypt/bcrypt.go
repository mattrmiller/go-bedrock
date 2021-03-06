// Package brcrypt provides useful encryption and hashing functions.
package brcrypt

// Imports
import (
	"golang.org/x/crypto/bcrypt"
)

// BcryptHash Bcrypt hash with defined number of cycles.
func BcryptHash(value string, cycles int) (string, error) {

	// Hash
	ret, err := bcrypt.GenerateFromPassword([]byte(value), cycles)
	if err != nil {
		return "", err
	}

	return string(ret), nil
}

// MustBcryptHash Bcrypt hash with defined number of cycles. Panics on error.
func MustBcryptHash(value string, cycles int) string {

	// Hash
	ret, err := BcryptHash(value, cycles)
	if err != nil {
		panic(err)
	}

	return ret
}

// BcryptPasswordCompare Bcrypt hash compare.
func BcryptPasswordCompare(hashA string, hashB string) error {

	// Compare hashes
	err := bcrypt.CompareHashAndPassword([]byte(hashA), []byte(hashB))
	if err != nil {
		return err
	}

	return nil
}
