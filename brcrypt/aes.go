// Package brcrypt provides useful encryption and hashing functions.
package brcrypt

// Imports
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"

	"github.com/mattrmiller/go-bedrock/brstrings"
)

// Constants
const (
	AesKeyLength128 = 16
	AesKeyLength192 = 24
	AesKeyLength256 = 32
)

// AesKey128 Generate key for AES-128.
func AesKey128() ([]byte, error) {

	// Generate random bytes
	return brstrings.RandomBytes(AesKeyLength128)
}

// AesKey192 Generate key for AES-192.
func AesKey192() ([]byte, error) {

	// Generate randome bytes
	return brstrings.RandomBytes(AesKeyLength192)
}

// AesKey256 Generate key for AES-256.
func AesKey256() ([]byte, error) {

	// Generate random bytes.
	return brstrings.RandomBytes(AesKeyLength256)
}

// AesEncrypt Aes encrypt in GCM mode value with key, depending on key length determines
// the AES algorithm that will be used.
func AesEncrypt(value string, key []byte) (string, error) {

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Seal
	encrypted := gcm.Seal(nonce, nonce, []byte(value), nil)

	// Format in hex
	return fmt.Sprintf("%x", encrypted), nil
}

// MustAesEncrypt Aes encrypt in GCM mode value with key, depending on key length determines
// the AES algorithm that will be used. Panics on error.
func MustAesEncrypt(value string, key []byte) string {

	// Encrypt
	ret, err := AesEncrypt(value, key)
	if err != nil {
		panic(err)
	}

	return ret
}

// AesDecrypt Aes decrypt in GCM mode value with key, depending on key length determines
// the AES algorithm that will be used.
func AesDecrypt(value string, key []byte) (string, error) {

	// Cipher text
	decoded, err := hex.DecodeString(value)
	if err != nil {
		return "", err
	}
	cipherValue := []byte(decoded)

	// Create cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	if len(cipherValue) < nonceSize {
		return "", errors.New("Ciphered value is too short")
	}

	// Unseal
	nonce, cipherValue := cipherValue[:nonceSize], cipherValue[nonceSize:]
	unCipherValue, err := gcm.Open(nil, nonce, cipherValue, nil)
	if err != nil {
		return "", nil
	}

	// Format in hex
	return string(unCipherValue), nil
}

// MustAesDecrypt Aes decrypt in GCM mode value with key, depending on key length determines
// the AES algorithm that will be used. Panics on error.
func MustAesDecrypt(value string, key []byte) string {

	// Decrypt
	ret, err := AesDecrypt(value, key)
	if err != nil {
		panic(err)
	}

	return ret
}
