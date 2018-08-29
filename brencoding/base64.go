// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"encoding/base64"
)

// Base64Encode Base64 encode, as defined in RFC 4648.
func Base64Encode(value string) string {

	// Encode
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// Base64Decode Base64 decode, as defined in RFC 4648.
func Base64Decode(value string) (string, error) {

	// Decode
	ret, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	// Return
	return string(ret), nil
}

// MustBase64Decode Base64 decode, as defined in RFC 4648. Panics on error.
func MustBase64Decode(value string) string {

	// Decode
	ret, err := Base64Decode(value)
	if err != nil {
		panic(err)
	}

	// Return
	return ret
}

// Base64UrlEncode Base64 encode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names.
func Base64UrlEncode(value string) string {

	// Encode
	return base64.URLEncoding.EncodeToString([]byte(value))
}

// Base64UrlDecode Base64 decode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names.
func Base64UrlDecode(value string) (string, error) {

	// Decode
	ret, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	// Return
	return string(ret), nil
}

// MustBase64UrlDecode Base64 decode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names. Panics on error.
func MustBase64UrlDecode(value string) string {

	// Decode
	ret, err := Base64UrlDecode(value)
	if err != nil {
		panic(err)
	}

	// Return
	return ret
}

// Base64EncodeNoPadding Base64 encode, as defined in RFC 4648. This is the same as standard encoding
// but omits padding characters.
func Base64EncodeNoPadding(value string) string {

	// Encode
	return base64.RawStdEncoding.EncodeToString([]byte(value))
}

// Base64DecodeNoPadding Base64 decode, as defined in RFC 4648. This is the same as standard encoding
// but omits padding characters.
func Base64DecodeNoPadding(value string) (string, error) {

	// Decode
	ret, err := base64.RawStdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	// Return
	return string(ret), nil
}

// MustBase64DecodeNoPadding Base64 decode, as defined in RFC 4648. This is the same as standard encoding
// but omits padding characters. Panics on error.
func MustBase64DecodeNoPadding(value string) string {

	// Decode
	ret, err := Base64DecodeNoPadding(value)
	if err != nil {
		panic(err)
	}

	// Return
	return ret
}

// Base64UrlEncodeNoPadding Base64 encode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names. This is the same as
// url encoding but omits padding characters.
func Base64UrlEncodeNoPadding(value string) string {

	// Encode
	return base64.RawURLEncoding.EncodeToString([]byte(value))
}

// Base64UrlDecodeNoPadding Base64 decode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names. This is the same as
// url encoding but omits padding characters.
func Base64UrlDecodeNoPadding(value string) (string, error) {

	// Decode
	ret, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	// Return
	return string(ret), nil
}

// MustBase64UrlDecodeNoPadding Base64 decode, using URLEncoding is the alternate base64 encoding defined in
// RFC 4648. It is typically used in URLs and file names. This is the same as
// url encoding but omits padding characters. Panics on error.
func MustBase64UrlDecodeNoPadding(value string) string {

	// Decode
	ret, err := Base64UrlDecodeNoPadding(value)
	if err != nil {
		panic(err)
	}

	// Return
	return ret
}
