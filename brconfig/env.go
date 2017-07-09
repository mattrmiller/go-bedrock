// Package brconfig provides useful configuration management functionality.
package brconfig

// Imports
import (
	"os"
	"strings"

	"github.com/mattrmiller/go-bedrock/brstrings"
)

// Environment as string.
func EnvAsString(key string) string {
	return os.Getenv(key)
}

// Environment as bool.
func EnvAsBool(key string) (bool, error) {

	// Get value
	value := EnvAsString(key)

	// Parse float
	ret, err := brstrings.StringToBool(value)
	if err != nil {
		return false, err
	}

	return ret, err
}

// Environment as bool. Panics on error.
func MustEnvAsBool(key string) bool {

	// Get value
	value, err := EnvAsBool(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as Int.
func EnvAsInt(key string) (int, error) {

	// Get value
	value := EnvAsString(key)

	// Convert to Int
	ret, err := brstrings.StringToInt(value)
	if err != nil {
		return int(0), err
	}

	return ret, nil
}

// Environment as Int. Panics on error.
func MustEnvAsInt(key string) int {

	// Get value
	value, err := EnvAsInt(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as Int64.
func EnvAsInt64(key string) (int64, error) {

	// Get value
	value := EnvAsString(key)

	// Parse int
	ret, err := brstrings.StringToInt64(value)
	if err != nil {
		return int64(0), err
	}

	return ret, nil
}

// Environment as Int64. Panics on error.
func MustEnvAsInt64(key string) int64 {

	// Get value
	value, err := EnvAsInt64(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as float.
func EnvAsFloat(key string) (float64, error) {

	// Get value
	value := EnvAsString(key)

	// Parse float
	ret, err := brstrings.StringToFloat(value)
	if err != nil {
		return float64(0), err
	}

	return ret, nil
}

// Environment as float. Panics on error.
func MustEnvAsFloat(key string) float64 {

	// Get value
	value, err := EnvAsFloat(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment variable as slice, defining your own separator to use.
func EnvAsSlice(key, separator string) []string {

	// Get value
	value := EnvAsString(key)

	// Split
	return strings.Split(value, separator)
}
