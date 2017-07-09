// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"encoding/json"
)

// Encode Json.
func EncodeJson(payload interface{}) (string, error) {

	// Marshal
	ret, err := json.Marshal(payload)
	if err != nil {
		return "", nil
	}

	return string(ret), nil
}

// Encode Json. Panics on error.
func MustEncodeJson(payload interface{}) string {

	// Marshal
	ret, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return string(ret)
}

// Encode Json in a pretty format with specified indenting.
func EncodeJsonPretty(payload interface{}, indent string) (string, error) {

	// Marshal
	ret, err := json.MarshalIndent(payload, "", indent)
	if err != nil {
		return "", nil
	}

	return string(ret), nil
}

// Encode Json in a pretty format with specified indenting.
func MustEncodeJsonPretty(payload interface{}, indent string) string {

	// Marshal
	ret, err := json.MarshalIndent(payload, "", indent)
	if err != nil {
		panic(err)
	}

	return string(ret)
}

// Decode Json.
func DecodeJson(payload string) (map[string]interface{}, error) {

	// Return
	var ret interface{}

	// Unmarshal
	err := json.Unmarshal([]byte(payload), &ret)
	if err != nil {
		return nil, err
	}

	return ret.(map[string]interface{}), nil
}

// Decode Json. Panics on error.
func MustDecodeJson(payload string) map[string]interface{} {

	// Return
	var ret interface{}

	// Unmarshal
	err := json.Unmarshal([]byte(payload), &ret)
	if err != nil {
		panic(err)
	}

	return ret.(map[string]interface{})
}
