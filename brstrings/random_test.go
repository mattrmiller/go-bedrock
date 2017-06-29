// Package brstrings provides useful string functions
package brstrings

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test randoms
func TestRandoms(tst *testing.T) {

	// Test random alpha numeric
	r1 := RandomAlphaNumString(20)
	r2 := RandomAlphaNumString(20)
	brtesting.AssertNotEqual(tst, r1, r2, "RandomAlphaNumString failed")

}
