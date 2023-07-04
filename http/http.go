package http

import (
	"encoding/json"
	"net/http"
)

// EncodeResponse encodes the response to JSON and writes it to the response writer
func EncodeResponse(w http.ResponseWriter, response interface{}) {
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.Write([]byte("Error encoding response"))
	}
}

// EncodeError encodes the error to JSON and writes it to the response writer
func EncodeError(w http.ResponseWriter, response error) {
	w.WriteHeader(http.StatusInternalServerError)
	EncodeResponse(w, response.Error())
}
