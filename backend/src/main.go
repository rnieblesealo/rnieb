package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gographics/imagick.v3/imagick"
	"heicupload.rnieb.dev/heicupload"
)

const (
	PORT = ":8080"
)

func main() {
	// initialize imagemagick
	imagick.Initialize()
	defer imagick.Terminate()

	// setup handlers
	http.HandleFunc("/ping", heicupload.PingHandler)
	http.HandleFunc("/upload", heicupload.UploadHandler)

	fmt.Printf("Launching server on port %s...\n", PORT)

	log.Fatal(http.ListenAndServe(PORT, nil)) // listenandserve always returns non nil err
}
