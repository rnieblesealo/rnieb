package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"rnieb/auth"
	"rnieb/fetch"
	"rnieb/upload"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	PORT = ":8080"
)

// Allows any origin to access this; effectively we're a public API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// allow get, post delete only
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")

		// TODO: not sure what this does
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

	// Open database connection

	db, err := sql.Open("sqlite3", "./rnieb.db") // Open DB connection
	if err != nil {
		log.Fatalf("Failed to open DB connection: %s\n", err)
	}
	defer db.Close()

	// Setup auth handlers
	http.HandleFunc("/login", auth.LoginHandler(db))

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
