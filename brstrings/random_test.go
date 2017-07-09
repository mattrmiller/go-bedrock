// Package brstrings provides useful string functions.
package brstrings

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test random strings.
func TestRandomStrings(tst *testing.T) {

	// Test random alpha numeric
	r1 := RandomAlphaNumString(20)
	r2 := RandomAlphaNumString(20)
	brtesting.AssertNotEqual(tst, r1, r2, "RandomAlphaNumString failed")
}

// Test random bytes.
func TestRandomBytes(tst *testing.T) {

	// Test random bytes
	r1, err1 := RandomBytes(20)
	r2, err2 := RandomBytes(20)
	brtesting.AssertEqual(tst, err1, nil, "RandomBytes failed")
	brtesting.AssertEqual(tst, err2, nil, "RandomBytes failed")
	brtesting.AssertNotEqual(tst, string(r1), string(r2), "RandomBytes failed")
}
