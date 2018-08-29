// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"encoding/json"
)

// EncodeJSON Encode JSON payload.
func EncodeJSON(payload interface{}) (string, error) {

	// Marshal
	ret, err := json.Marshal(payload)
	if err != nil {
		return "", nil
	}

	return string(ret), nil
}

// MustEncodeJSON Encode JSON payload. Panics on error.
func MustEncodeJSON(payload interface{}) string {

	// Marshal
	ret, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return string(ret)
}

// EncodeJSONPretty Encode JSON payload in a pretty format with specified indenting.
func EncodeJSONPretty(payload interface{}, indent string) (string, error) {

	// Marshal
	ret, err := json.MarshalIndent(payload, "", indent)
	if err != nil {
		return "", nil
	}

	return string(ret), nil
}

// MustEncodeJSONPretty Encode JSON payload in a pretty format with specified indenting.
func MustEncodeJSONPretty(payload interface{}, indent string) string {

	// Marshal
	ret, err := json.MarshalIndent(payload, "", indent)
	if err != nil {
		panic(err)
	}

	return string(ret)
}

// DecodeJSON Decode JSON payload.
func DecodeJSON(payload string) (map[string]interface{}, error) {

	// Return
	var ret interface{}

	// Unmarshal
	err := json.Unmarshal([]byte(payload), &ret)
	if err != nil {
		return nil, err
	}

	return ret.(map[string]interface{}), nil
}

// MustDecodeJSON Decode JSON payload. Panics on error.
func MustDecodeJSON(payload string) map[string]interface{} {

	// Return
	var ret interface{}

	// Unmarshal
	err := json.Unmarshal([]byte(payload), &ret)
	if err != nil {
		panic(err)
	}

	return ret.(map[string]interface{})
}
