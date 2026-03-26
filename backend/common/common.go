package common

import (
	"encoding/json"
	"net/http"
)

// Standard constants
const (
	DB_PHOTO_TYPE_STR = "photo"
	DB_VIDEO_TYPE_STR = "video"
	DB_AUDIO_TYPE_STR = "audio"
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

// Returns allowed DB media types as a tuple
// Useful for MatchesAny checks
func GetAllowedDBMediaTypes() []string {
	return []string{
		DB_PHOTO_TYPE_STR,
		DB_VIDEO_TYPE_STR,
		DB_AUDIO_TYPE_STR,
	}
}
