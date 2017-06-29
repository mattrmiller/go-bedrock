// Package brcrypt provides useful encryption and hashing functions
package brcrypt

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test Sha hashing
func TestShaHashing(tst *testing.T) {

	// Test sha
	brtesting.AssertEqual(tst, HashSha512("go-bedrock"), "4b15413e6ca54313e2c86ca7c4f20da74db6083203b3b75ba73a0b9f10068f84c2ed77a268f86bb231259595a25be3b2b29529d4f544bc9452676e185da83c6a", "HashSha512 failed")
	brtesting.AssertEqual(tst, HashSha384("go-bedrock"), "87542522d90176c8eeca6d348940483ea055523cc3d30500f1586e8dd9145efb9e93180867d440ac131e9b8cf09a946a", "HashSha384 failed")
	brtesting.AssertEqual(tst, HashSha256("go-bedrock"), "7a52b82a3ab1ecfc0fc9d0e57cabe6278063905aa973a28e37cb670d33bebaca", "HashSha256 failed")
	brtesting.AssertEqual(tst, HashSha1("go-bedrock"), "a72d9b86ab715924b1a8d68c857ced3e9cf54eae", "HashSha1 failed")
}

// Test Md5 hashing
func TestMd5Hashing(tst *testing.T) {

	// Test md5
	brtesting.AssertEqual(tst, HashMd5("go-bedrock"), "75a7fad8843bb0a4dc9496719338934d", "HashMd5 failed")
}
