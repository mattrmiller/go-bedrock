// Package brcrypt provides useful encryption and hashing functions
package brcrypt

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test bcrypt hashing
func TestBcryptHashing(tst *testing.T) {

	// Test bcrypt
	hashA, errA := BcryptHash("go-bedrock-secret-value-A", 12)
	hashB, errB := BcryptHash("go-bedrock-secret-value-B", 12)
	hashC, errC := BcryptHash("go-bedrock-secret-value-C", 12)
	brtesting.AssertEqual(tst, errA, nil, "BcryptHash failed")
	brtesting.AssertEqual(tst, errB, nil, "BcryptHash failed")
	brtesting.AssertEqual(tst, errC, nil, "BcryptHash failed")
	brtesting.AssertNotEqual(tst, hashA, hashB, "BcryptHash failed")
	brtesting.AssertNotEqual(tst, hashB, hashC, "BcryptHash failed")
	brtesting.AssertNotEqual(tst, hashA, hashC, "BcryptHash failed")
}
