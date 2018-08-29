// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test JSON encode.
func TestJSONEncode(tst *testing.T) {

	// Test encode
	payload := map[string]string{
		"hello": "world",
	}
	encoded, err := EncodeJSON(payload)
	brtesting.AssertEqual(tst, encoded, `{"hello":"world"}`, "EncodeJSON failed")
	brtesting.AssertEqual(tst, err, nil, "EncodeJSON failed")

	// Test encode pretty
	encoded, err = EncodeJSONPretty(payload, "      ")
	brtesting.AssertNotEqual(tst, encoded, `{"hello":"world"}`, "EncodeJSONPretty failed")
	brtesting.AssertEqual(tst, err, nil, "EncodeJSONPretty failed")
}

// Test JSON decode.
func TestJSONDecode(tst *testing.T) {

	// Test encode
	payload, err := DecodeJSON(`{"hello":"world"}`)
	brtesting.AssertEqual(tst, payload["hello"].(string), "world", "DecodeJSON failed")
	brtesting.AssertEqual(tst, err, nil, "DecodeJSON failed")

}
