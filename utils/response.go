package utils

import (
	"encoding/json"
	"net/http"
)

func WriteResponseToJson[T any](w http.ResponseWriter, statusCode int, v T) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

func WriteErrorToJson(w http.ResponseWriter, statusCode int, err error) error {
	errResponse := map[string]any{
		"success": false,
		"message": err.Error(),
	}

	return WriteResponseToJson(w, statusCode, errResponse)
}

func WriteSuccessToJson(w http.ResponseWriter, statusCode int, message string, v any) error {
	errResponse := map[string]any{
		"success": true,
		"message": message,
		"data":    v,
	}

	return WriteResponseToJson(w, statusCode, errResponse)
}
