// Package brstrings provides useful testing functions
package brtesting

// Imports
import (
	"testing"
)

// Test asserts
func TestAsserts(tst *testing.T) {

	// Assert equal
	AssertEqual(tst, 1, 1, "AssertEqual failed")

	// Assert not equal
	AssertNotEqual(tst, 1, 2, "AssertNotEqual failed")
}
