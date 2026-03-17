package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gographics/imagick.v3/imagick"
	"heicupload.rnieb.dev/heicupload"
	"heicupload.rnieb.dev/imgfetch"
)

const (
	PORT = ":8080"
)

func main() {
	// initialize imagemagick
	imagick.Initialize()
	defer imagick.Terminate()

	// setup image upload handlers
	http.HandleFunc("/ping", heicupload.PingHandler)
	http.HandleFunc("/upload", heicupload.UploadHandler)

	// setup image fetch handlers
	http.HandleFunc("/list-images", imgfetch.ListImages)

	/* set up a fileserver that will look for an image filename inside /uploads
	e.g. a request to /uploads/img0.png will extract the filename from url
	then look for that only in the fileserver */

	http.Handle("/uploads", // handle uploads route
		http.StripPrefix("/uploads/", // strip this prefix from url ( leaves only filename )
			http.FileServer(http.Dir("/uploads")))) // a fileserver for the /uploads dir

	fmt.Printf("Launching server on port %s...\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, nil)) // listenandserve always returns non nil err
}
