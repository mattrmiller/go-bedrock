// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test Base64.
func TestBase64(tst *testing.T) {

	// Test Base64
	encode := Base64Encode("go-bedrock-base64")
	brtesting.AssertEqual(tst, encode, "Z28tYmVkcm9jay1iYXNlNjQ=", "Base64Encode failed")
	decoded, err := Base64Decode(encode)
	brtesting.AssertEqual(tst, decoded, "go-bedrock-base64", "Base64Decode failed")
	brtesting.AssertEqual(tst, err, nil, "Base64Decode failed")
}

// Test Base64 with Url encoding.
func TestBase64Url(tst *testing.T) {

	// Test Base64 Url
	encode := Base64UrlEncode("go-bedrock-base64url")
	print(encode)
	brtesting.AssertEqual(tst, encode, "Z28tYmVkcm9jay1iYXNlNjR1cmw=", "Base64UrlEncode failed")
	decoded, err := Base64UrlDecode(encode)
	brtesting.AssertEqual(tst, decoded, "go-bedrock-base64url", "Base64UrlDecode failed")
	brtesting.AssertEqual(tst, err, nil, "Base64Decode failed")
}

// Test Base64 with no padding.
func TestBase64NoPadding(tst *testing.T) {

	// Test Base64 raw
	encode := Base64EncodeNoPadding("go-bedrock-base64nopadding")
	brtesting.AssertEqual(tst, encode, "Z28tYmVkcm9jay1iYXNlNjRub3BhZGRpbmc", "Base64EncodeNoPadding failed")
	decoded, err := Base64DecodeNoPadding(encode)
	brtesting.AssertEqual(tst, decoded, "go-bedrock-base64nopadding", "Base64DecodeNoPadding failed")
	brtesting.AssertEqual(tst, err, nil, "Base64Decode failed")
}

// Test Base64 with Url encoding and no padding.
func TestBase64UrlNoPadding(tst *testing.T) {

	// Test Base64 Url
	encode := Base64UrlEncodeNoPadding("go-bedrock-base64urlnopadding")
	brtesting.AssertEqual(tst, encode, "Z28tYmVkcm9jay1iYXNlNjR1cmxub3BhZGRpbmc", "Base64UrlEncode failed")
	decoded, err := Base64UrlDecodeNoPadding(encode)
	brtesting.AssertEqual(tst, decoded, "go-bedrock-base64urlnopadding", "Base64UrlDecode failed")
	brtesting.AssertEqual(tst, err, nil, "Base64Decode failed")
}
