// Package brencoding provides useful encoding functions.
package brencoding

// Imports
import (
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test Json encode.
func TestJsonEncode(tst *testing.T) {

	// Test encode
	payload := map[string]string{
		"hello": "world",
	}
	encoded, err := EncodeJson(payload)
	brtesting.AssertEqual(tst, encoded, `{"hello":"world"}`, "EncodeJson failed")
	brtesting.AssertEqual(tst, err, nil, "EncodeJson failed")

	// Test encode pretty
	encoded, err = EncodeJsonPretty(payload, "      ")
	brtesting.AssertNotEqual(tst, encoded, `{"hello":"world"}`, "EncodeJsonPretty failed")
	brtesting.AssertEqual(tst, err, nil, "EncodeJsonPretty failed")
}

// Test Json decode.
func TestJsonDecode(tst *testing.T) {

	// Test encode
	payload, err := DecodeJson(`{"hello":"world"}`)
	brtesting.AssertEqual(tst, payload["hello"].(string), "world", "DecodeJson failed")
	brtesting.AssertEqual(tst, err, nil, "DecodeJson failed")

}
