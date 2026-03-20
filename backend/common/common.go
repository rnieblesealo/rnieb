package common

import (
	"encoding/json"
	"net/http"
)

// Standard HTTP response
type RNResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Standard HTTP response function
func RNRespond(w http.ResponseWriter, message string, data interface{}, httpStatus int) {
	resp := RNResponse{
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	json.NewEncoder(w).Encode(resp)
}
