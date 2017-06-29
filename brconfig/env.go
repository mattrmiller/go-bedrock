// Package brconfig provides useful configuration management functionality
package brconfig

// Imports
import (
	"os"
	"strconv"
	"strings"
)

// Environment as string
func EnvAsString(key string) string {
	return os.Getenv(key)
}

// Environment as bool
func EnvAsBool(key string) (bool, error) {

	// Get value
	value := EnvAsString(key)

	// Parse float
	ret, err := strconv.ParseBool(value)
	if err != nil {
		return false, err
	}

	return ret, err
}

// Environment as bool, with panic on error
func EnvAsBoolPanic(key string) bool {

	// Get value
	value, err := EnvAsBool(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as float
func EnvAsFloat(key string) (float64, error) {

	// Get value
	value := EnvAsString(key)

	// Parse float
	ret, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return float64(0), err
	}

	return ret, nil
}

// Environment as float, with panic on error
func EnvAsFloatPanic(key string) float64 {

	// Get value
	value, err := EnvAsFloat(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as Int
func EnvAsInt(key string) (int, error) {

	// Get value
	value := EnvAsString(key)

	// Convert to Int
	ret, err := strconv.Atoi(value)
	if err != nil {
		return int(0), err
	}

	return ret, nil
}

// Environment as Int, with panic on error
func EnvAsIntPanic(key string) int {

	// Get value
	value, err := EnvAsInt(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment as Int64
func EnvAsInt64(key string) (int64, error) {

	// Get value
	value := EnvAsString(key)

	// Parse float
	ret, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return int64(0), err
	}

	return ret, nil
}

// Environment as Int64, with panic on error
func EnvAsInt64Panic(key string) int64 {

	// Get value
	value, err := EnvAsInt64(key)
	if err != nil {
		panic(err.Error())
	}

	return value
}

// Environment variable as slice, defining your own separator to use
func EnvAsSlice(key, separator string) []string {

	// Get value
	value := EnvAsString(key)

	// Split
	return strings.Split(value, separator)
}
