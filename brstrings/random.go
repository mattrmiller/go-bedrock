// Package brstrings provides useful string functions
package brstrings

// Imports
import (
	"math/rand"
	"time"
)

// Random string of defined length, defining your list of characters to use
func RandomString(length int, chars string) string {

	// Seed random
	rand.Seed(time.Now().UTC().UnixNano())

	// Make result
	result := make([]byte, length)

	// Generate
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}

// Random alpha numeric string, of defined length
func RandomAlphaNumString(length int) string {

	// Random string, with alpha numeric
	return RandomString(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
}

// Random bytes, of defined length
func RandomBytes(length int) ([]byte, error) {

	// Make result
	result := make([]byte, length)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	_, err := r.Read(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
