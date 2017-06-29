// Package brconfig provides useful configuration management functionality
package brconfig

// Imports
import (
	"os"
	"testing"

	"github.com/mattrmiller/go-bedrock/brtesting"
)

// Test env values
func TestEnvValues(tst *testing.T) {

	// Set some values
	os.Setenv("TEST_STRING_HELLO", "hello")
	os.Setenv("TEST_BOOL_TRUE", "true")
	os.Setenv("TEST_BOOL_FALSE", "false")
	os.Setenv("TEST_BOOL_INVALID", "invalid value")
	os.Setenv("TEST_FLOAT_1345", "1.345")
	os.Setenv("TEST_FLOAT_INVALID", "invalid value")
	os.Setenv("TEST_INT_80", "80")
	os.Setenv("TEST_INT_INVALID", "invalid value")
	os.Setenv("TEST_INT64_100", "100")
	os.Setenv("TEST_INT64_INVALID", "invalid value")

	// Test string
	s := EnvAsString("TEST_STRING_HELLO")
	brtesting.AssertEqual(tst, s, "hello", "EnvAsString failed")

	// Test bool
	b, err := EnvAsBool("TEST_BOOL_TRUE")
	brtesting.AssertEqual(tst, b, true, "EnvAsBool failed")
	brtesting.AssertEqual(tst, err, nil, "EnvAsBool failed")
	b, err = EnvAsBool("TEST_BOOL_FALSE")
	brtesting.AssertEqual(tst, b, false, "EnvAsBool failed")
	brtesting.AssertEqual(tst, err, nil, "EnvAsBool failed")
	b, err = EnvAsBool("TEST_BOOL_INVALID")
	brtesting.AssertEqual(tst, b, false, "EnvAsBool failed")
	brtesting.AssertNotEqual(tst, err, nil, "EnvAsBool failed")

	// Test float
	f, err := EnvAsFloat("TEST_FLOAT_1345")
	brtesting.AssertEqual(tst, f, float64(1.345), "EnvAsFloat failed")
	brtesting.AssertEqual(tst, err, nil, "EnvAsFloat failed")
	f, err = EnvAsFloat("TEST_FLOAT_INVALID")
	brtesting.AssertEqual(tst, f, float64(0), "EnvAsFloat failed")
	brtesting.AssertNotEqual(tst, err, nil, "EnvAsFloat failed")

	// Test int
	i, err := EnvAsInt("TEST_INT_80")
	brtesting.AssertEqual(tst, i, 80, "EnvAsInt failed")
	brtesting.AssertEqual(tst, err, nil, "EnvAsInt failed")
	i, err = EnvAsInt("TEST_INT_INVALID")
	brtesting.AssertEqual(tst, i, 0, "EnvAsInt failed")
	brtesting.AssertNotEqual(tst, err, nil, "EnvAsInt failed")

	// Test in64
	is, err := EnvAsInt64("TEST_INT64_100")
	brtesting.AssertEqual(tst, is, int64(100), "EnvAsInt64 failed")
	brtesting.AssertEqual(tst, err, nil, "EnvAsInt64 failed")
	is, err = EnvAsInt64("TEST_INT64_INVALID")
	brtesting.AssertEqual(tst, is, int64(0), "EnvAsInt64 failed")
	brtesting.AssertNotEqual(tst, err, nil, "EnvAsInt64 failed")

}
