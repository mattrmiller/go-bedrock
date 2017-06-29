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
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}

	return string(result)
}

// Random alpha numeric string, of defined length
func RandomAlphaNumString(length int) string {

	// Random string, with alpha numeric
	return RandomString(length, "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
}
