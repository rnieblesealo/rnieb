package imgfetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// TODO: structure the responses similarly to other http

func ListImages(w http.ResponseWriter, r *http.Request) {
	dirEntries, err := os.ReadDir("/uploads")
	if err != nil {
		fmt.Printf("ListImages: %s\n", err.Error())
		return
	}

	var names []string
	for _, dirEntry := range dirEntries {
		names = append(names, dirEntry.Name())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(names)
}
