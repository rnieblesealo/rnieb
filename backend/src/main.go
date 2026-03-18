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

// cors middleware
// this assumes our program is a fully public api
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // fix cors
		next.ServeHTTP(w, r)
	})

}

func main() {
	// initialize imagemagick
	imagick.Initialize()
	defer imagick.Terminate()

	// setup image upload handlers
	http.HandleFunc("/ping", upload.Ping)
	http.HandleFunc("/upload", upload.Upload)

	// setup image fetch handlers
	http.HandleFunc("/list-images", fetch.ListImages) // WARNING: use get-drawings!

	http.HandleFunc("/get-drawings", fetch.GetDrawings)

	/* set up a fileserver that will look for an image filename inside /uploads
	e.g. a request to /uploads/img0.png will extract the filename from url
	then look for that only in the fileserver */

	http.Handle("/uploads/", // handle uploads route
		http.StripPrefix("/uploads/", // strip this prefix from url ( leaves only filename )
			http.FileServer(http.Dir("/uploads")))) // a fileserver for the /uploads dir

	fmt.Printf("Launching server on port %s...\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, corsMiddleware(http.DefaultServeMux))) // listenandserve always returns non nil err
	// add cors middleware when serving!
}
