package main

import (
	"encoding/json"
	"net/http"
)

func ParseJSONBody(r *http.Request, dst interface{}) error {
	// Decode the request body into the destination
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
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
