package upload

import (
	"database/sql"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"rnieb/common"
	"time"
)

const (
	UPLOAD_DIR = "/uploads"

	FORM_IMAGE_FILE_NAME        = "file"
	FORM_IMAGE_DESCRIPTION_NAME = "description"
	FORM_IMAGE_NAME_NAME        = "name"

	MAX_FORM_SIZE = 32 << 20 // 32mb upload limit; converting to mebi with shift!
)

// Check for heartbeat
func Ping(w http.ResponseWriter, req *http.Request) {
	common.RNRespond(w, "Polo!", nil, http.StatusOK)
}

// Handles image uploads
func Upload(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Parse request form

		err := req.ParseMultipartForm(MAX_FORM_SIZE)
		if err != nil {
			common.RNRespond(w, fmt.Sprintf("Failed to parse form: %s", err), nil, http.StatusBadRequest)
			return
		}

		// Get request form values

		imageName := req.FormValue("name")
		imageDescription := req.FormValue("description")

		_, imageFileHandle, err := req.FormFile("file")
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to obtain image file: %s", err),
				nil,
				http.StatusBadRequest)

			return
		}

		// Open image file

		file, err := imageFileHandle.Open()
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to open image file: %s", err),
				nil,
				http.StatusBadRequest)

			return
		}

		defer file.Close()

		// Read first 512 bytes of the file to obtain MIME type

		buf := make([]byte, 512)
		_, err = file.Read(buf)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to read image file: %s", err),
				nil,
				http.StatusBadRequest)

			return
		}

		// Ensure image upload is PNG, JPEG or HEIC
		/* (Might support more later but these are most common ones!) */

		allowedMimeTypes := []string{
			"image/png",
			"image/jpeg",
			"image/heic",
		}

		mimeType := mimetype.Detect(buf).String()
		if !mimetype.EqualsAny(mimeType, allowedMimeTypes...) {
			common.RNRespond(w, "Uploaded image must be PNG, JPG or HEIC", nil, http.StatusBadRequest)
			return
		}

		// Move read cursor to beginning of file

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to seek: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Create uploads dir

		err = os.MkdirAll(UPLOAD_DIR, os.ModePerm)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to create uploads dir: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Create the image file
		// Will be named the current unix nano time to uniqify

		imageFilepath :=
			fmt.Sprintf("%s/%d%s",
				UPLOAD_DIR,
				time.Now().UnixNano(),
				filepath.Ext(imageFileHandle.Filename))

		imageFile, err := os.Create(imageFilepath)

		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to create image file: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}
		defer imageFile.Close()

		// Copy uploaded image data to image file

		_, err = io.Copy(imageFile, file)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to write image file contents: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Ensure image is PNG and of width 500
		/* Want everything to be in same format for standardization;
		PNG is ideal for web;
		width 500 keeps things small but still clear enough for our purposes */

		var pngImageFilepath string

		if mimeType != "image/png" {
			// If did not already upload PNG, perform conversion
			pngImageFilepath, err = ConvertToPNG(imageFile.Name())
			if err != nil {
				common.RNRespond(
					w,
					fmt.Sprintf("Failed to convert image to PNG: %s", err),
					nil,
					http.StatusInternalServerError)

				return
			}

			pngImageFilepath, _ = ResizePNG(pngImageFilepath)
		} else {
			// Only resize if already PNG
			pngImageFilepath, _ = ResizePNG(imageFile.Name())
		}

		// Create image entry in DB

		_, err = db.Exec(`
			INSERT INTO drawings (name, description, path)
			VALUES (?, ?, ?)	
	`, imageName, imageDescription, pngImageFilepath) // Run insertion query
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to insert image into DB: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Respond successfully

		common.RNRespond(
			w,
			fmt.Sprintf("Successfully uploaded and processed %s", pngImageFilepath),
			nil,
			http.StatusOK)
	}
}
