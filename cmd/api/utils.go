package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

// Validate is a pointer to a validator instance
var Validate *validator.Validate

// ParseJSONBody is a helper function that decodes the request body into the destination
func init() {
	Validate = validator.New()
}

func ParseJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Decode the request body into the destination
	// MaxBytesReader is used to prevent the request from reading too much data
	maxBytes := int64(1 << 20) // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	// Decode the request body into the destination
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(dst)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	// Set the content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Marshal the data into JSON and write it to the response
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}

	return nil
}

func WriteError(w http.ResponseWriter, status int, err error) {
	// Set the content type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Create a map to hold the error message
	payload := map[string]string{"error": err.Error()}

	// Marshal the data into JSON and write it to the response
	json.NewEncoder(w).Encode(payload)
}
