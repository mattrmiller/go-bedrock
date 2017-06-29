// Package brcrypt provides useful encryption and hashing functions
package brcrypt

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test AES-128
func TestAes128(tst *testing.T) {

	// Test AES-128
	keyA, errA := AesKey128()
	keyB, errB := AesKey128()
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, string(keyA), string(keyB), "AesEncrypt failed")
	valueA, errA := AesEncrypt("go-bedrock-secret-value-A", keyA)
	valueB, errB := AesEncrypt("go-bedrock-secret-value-B", keyB)
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, valueA, valueB, "AesEncrypt failed")
	valueC, errA := AesDecrypt(valueA, keyA)
	valueD, errB := AesDecrypt(valueB, keyB)
	brtesting.AssertEqual(tst, valueC, "go-bedrock-secret-value-A", "AesDecrypt failed")
	brtesting.AssertEqual(tst, valueD, "go-bedrock-secret-value-B", "AesDecrypt failed")
}

// Test AES-192
func TestAes192(tst *testing.T) {

	// Test AES-128
	keyA, errA := AesKey192()
	keyB, errB := AesKey192()
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, string(keyA), string(keyB), "AesEncrypt failed")
	valueA, errA := AesEncrypt("go-bedrock-secret-value-A", keyA)
	valueB, errB := AesEncrypt("go-bedrock-secret-value-B", keyB)
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, valueA, valueB, "AesEncrypt failed")
	valueC, errA := AesDecrypt(valueA, keyA)
	valueD, errB := AesDecrypt(valueB, keyB)
	brtesting.AssertEqual(tst, valueC, "go-bedrock-secret-value-A", "AesDecrypt failed")
	brtesting.AssertEqual(tst, valueD, "go-bedrock-secret-value-B", "AesDecrypt failed")
}

// Test AES-256
func TestAes256(tst *testing.T) {

	// Test AES-128
	keyA, errA := AesKey256()
	keyB, errB := AesKey256()
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, string(keyA), string(keyB), "AesEncrypt failed")
	valueA, errA := AesEncrypt("go-bedrock-secret-value-A", keyA)
	valueB, errB := AesEncrypt("go-bedrock-secret-value-B", keyB)
	brtesting.AssertEqual(tst, errA, nil, "AesEncrypt failed")
	brtesting.AssertEqual(tst, errB, nil, "AesEncrypt failed")
	brtesting.AssertNotEqual(tst, valueA, valueB, "AesEncrypt failed")
	valueC, errA := AesDecrypt(valueA, keyA)
	valueD, errB := AesDecrypt(valueB, keyB)
	brtesting.AssertEqual(tst, valueC, "go-bedrock-secret-value-A", "AesDecrypt failed")
	brtesting.AssertEqual(tst, valueD, "go-bedrock-secret-value-B", "AesDecrypt failed")
}
