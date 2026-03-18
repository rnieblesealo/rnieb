package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gographics/imagick.v3/imagick"
	"rnieb/fetch"
	"rnieb/upload"
)

const (
	PORT = ":8080"
)

// Allows any origin to access this; effectively we're a public API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	// Initialize ImageMagick
	/* Must be done once; C API */

	imagick.Initialize()
	defer imagick.Terminate()

	// Setup upload handlers

	http.HandleFunc("/ping", upload.Ping)
	http.HandleFunc("/upload", upload.Upload)

	// Setup fetch handlers

	http.HandleFunc("/get-drawings", fetch.GetDrawings)
	http.HandleFunc("/delete-drawing", fetch.DeleteDrawing)

	// Setup image fileserver

	http.Handle("/uploads/", // handle uploads route
		http.StripPrefix("/uploads/", // strip this prefix from url ( leaves only filename )
			http.FileServer(http.Dir("/uploads")))) // a fileserver for the /uploads dir

	fmt.Printf("Starting RNIEB server on port %s...\n", PORT)

	// Start HTTP server

	log.Fatal(http.ListenAndServe(PORT, corsMiddleware(http.DefaultServeMux)))
	// ListenAndServe always returns non nil err; only fails if it errors???
}
