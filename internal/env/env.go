package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// GetEnvVar is a generic function that returns the value of the environment variable key
func GetEnvVar[T any](key string, fallback T, parseFunc func(string) (T, error)) T {
	// Get the value of the environment variable key
	value, ok := os.LookupEnv(key)
	// If the environment variable key is not set, return the fallback value
	if !ok {
		return fallback
	}
	// Parse the value using the provided parse function
	parsedValue, err := parseFunc(value)
	if err != nil {
		return fallback
	}
	return parsedValue
}

// ParseString is a function that parses a string value to a string
func ParseString(value string) (string, error) {
	return value, nil
}

// ParseInt is a function that parses a string value to an integer
// - If the parsing fails, it returns an error
func ParseInt(value string) (int, error) {
	return strconv.Atoi(value)
}

// LoadEnv loads the environment variables from the .env file
func LoadEnv() {
	godotenv.Load()
}
