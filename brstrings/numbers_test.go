// Package brstrings provides useful string functions.
package brstrings

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test string to bool.
func TestBool(tst *testing.T) {

	// Test bool
	b, err := StringToBool("true")
	brtesting.AssertEqual(tst, err, nil, "StringToBool failed")
	brtesting.AssertEqual(tst, b, true, "StringToBool failed")
	b, err = StringToBool("True")
	brtesting.AssertEqual(tst, err, nil, "StringToBool failed")
	brtesting.AssertEqual(tst, b, true, "StringToBool failed")
	b, err = StringToBool("TruE")
	brtesting.AssertEqual(tst, err, nil, "StringToBool failed")
	brtesting.AssertEqual(tst, b, true, "StringToBool failed")
	b, err = StringToBool("false")
	brtesting.AssertEqual(tst, err, nil, "StringToBool failed")
	brtesting.AssertEqual(tst, b, false, "StringToBool failed")
	b, err = StringToBool("go-bedrock")
	brtesting.AssertNotEqual(tst, err, nil, "StringToBool failed")
}

// Test string to int.
func TestInt(tst *testing.T) {

	// Test bool
	i, err := StringToInt("80")
	brtesting.AssertEqual(tst, err, nil, "StringToInt failed")
	brtesting.AssertEqual(tst, i, 80, "StringToInt failed")
	i, err = StringToInt("-80")
	brtesting.AssertEqual(tst, err, nil, "StringToInt failed")
	brtesting.AssertEqual(tst, i, -80, "StringToInt failed")
	i, err = StringToInt("go-bedrock")
	brtesting.AssertNotEqual(tst, err, nil, "StringToInt failed")
}

// Test string to int64.
func TestInt64(tst *testing.T) {

	// Test bool
	i, err := StringToInt64("187480198367637651")
	brtesting.AssertEqual(tst, err, nil, "StringToInt64 failed")
	brtesting.AssertEqual(tst, i, int64(187480198367637651), "StringToInt64 failed")
	i, err = StringToInt64("go-bedrock")
	brtesting.AssertNotEqual(tst, err, nil, "StringToInt64 failed")
}

// Test string to float.
func TestFloat(tst *testing.T) {

	// Test bool
	f, err := StringToFloat("1.256898")
	brtesting.AssertEqual(tst, err, nil, "StringToFloat failed")
	brtesting.AssertEqual(tst, f, 1.256898, "StringToFloat failed")
	f, err = StringToFloat("go-bedrock")
	brtesting.AssertNotEqual(tst, err, nil, "StringToFloat failed")
}
