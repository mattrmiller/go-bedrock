// Package brcrypt provides useful encryption and hashing functions
package brcrypt

// Imports
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// Hash in sha512
func HashSha512(value string) string {

	// Sha
	sha := sha512.New()
	sha.Write([]byte(value))

	// Format in hex
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// Hash in sha384
func HashSha384(value string) string {

	// Sha
	sha := sha512.New384()
	sha.Write([]byte(value))

	// Format in hex
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// Hash in sha256
func HashSha256(value string) string {

	// Sha
	sha := sha256.New()
	sha.Write([]byte(value))

	// Format in hex
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// Hash in sha1
func HashSha1(value string) string {

	// Sha
	sha := sha1.New()
	sha.Write([]byte(value))

	// Format in hex
	return fmt.Sprintf("%x", sha.Sum(nil))
}

// Hash in md5
func HashMd5(value string) string {

	// Md5
	md5 := md5.New()
	md5.Write([]byte(value))

	// Format in hex
	return fmt.Sprintf("%x", md5.Sum(nil))
}
