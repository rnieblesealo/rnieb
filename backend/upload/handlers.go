package upload

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"rnieb/common"
	"time"

	"github.com/gabriel-vasile/mimetype"
	_ "github.com/mattn/go-sqlite3"
)

const (
	FORM_IMAGE_FILE_NAME        = "file"
	FORM_IMAGE_DESCRIPTION_NAME = "description"
	FORM_IMAGE_NAME_NAME        = "name"

	MAX_FORM_SIZE = 32 << 20 // 32mb upload limit; converting to mebi with shift!
)

// Check for heartbeat
func Ping(message string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		common.RNRespond(w, message, nil, http.StatusOK)
	}
}

// Handles media uploads
func UploadMedia(db *sql.DB, uploadPath string) http.HandlerFunc {

	// REQUIREMENTS:
	//
	//	1 file at a time packed as "file" in form data
	//	Type must be sent

	return func(w http.ResponseWriter, req *http.Request) {
		// Parse request form

		err := req.ParseMultipartForm(MAX_FORM_SIZE)
		if err != nil {
			common.RNRespond(w, fmt.Sprintf("Failed to parse form: %s", err), nil, http.StatusBadRequest)
			return
		}

		// Get request form values

		mediaName := req.FormValue("name")
		mediaDescription := req.FormValue("description")

		_, mediaFileHandle, err := req.FormFile("file")
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to obtain file handle: %s", err),
				nil,
				http.StatusBadRequest)

			return
		}

		// Open file

		file, err := mediaFileHandle.Open()
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to open file: %s", err),
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

		// Check if uploaded media type is valid

		allowedImageMimeTypes := []string{
			"image/png",
			"image/jpeg",
			"image/heic",
		}

		allowedVideoMimeTypes := []string{
			"video/quicktime",
		}

		allowed := append(allowedImageMimeTypes, allowedVideoMimeTypes...)

		mimeType := mimetype.Detect(buf).String()
		if !mimetype.EqualsAny(mimeType, allowed...) {
			common.RNRespond(
				w,
				"Uploaded file MIME type not recognized",
				nil,
				http.StatusBadRequest)

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

		// Ensure uploads dir

		err = os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to create uploads dir: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Create file in server
		// UNIX nano time name for uniqification

		mediaFilepath :=
			fmt.Sprintf("%s/%d%s",
				uploadPath,
				time.Now().UnixNano(),
				filepath.Ext(mediaFileHandle.Filename))
		mediaFile, err := os.Create(mediaFilepath)

		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to create media file: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}
		defer mediaFile.Close()

		// Copy contents from form file to media file

		_, err = io.Copy(mediaFile, file)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to write image file contents: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Set the media type to insert into DB
		// Do additional processing depending on MIME type
		// Update mediaFilepath to reflect conversions

		var dbMediaType string

		if mimetype.EqualsAny(mimeType, allowedImageMimeTypes...) {
			dbMediaType = common.DB_PHOTO_TYPE_STR

			// Ensure image is PNG and resize width
			//
			// Want everything to be in same format for standardization;
			// PNG is ideal for web
			// Width ~500 keeps things small but still clear enough for our purposes */

			if mimeType != "image/png" {
				// If did not already upload PNG, perform conversion
				mediaFilepath, err = ConvertToPNG(mediaFile.Name())
				if err != nil {
					common.RNRespond(
						w,
						fmt.Sprintf("Failed to convert image to PNG: %s", err),
						nil,
						http.StatusInternalServerError)

					return
				}

				// Then do resize
				mediaFilepath, _ = ResizePNG(mediaFilepath)
			} else {
				// Only resize if already PNG
				mediaFilepath, _ = ResizePNG(mediaFile.Name())
			}
		} else if mimetype.EqualsAny(mimeType, allowedVideoMimeTypes...) {
			dbMediaType = common.DB_VIDEO_TYPE_STR

			// Resize and compress all uploaded video

			mediaFilepath, err = ResizeAndCompressVideo(mediaFilepath)
			if err != nil {
				common.RNRespond(
					w,
					fmt.Sprintf("Failed to process video: %s", err),
					nil,
					http.StatusInternalServerError)

				return
			}
		}

		// Create entry in DB

		_, err = db.Exec(`
			insert into media (name, description, filename, type) values (?, ?, ?, ?)`,
			mediaName,
			mediaDescription,
			filepath.Base(mediaFilepath),
			dbMediaType)

		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to insert media into DB: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Respond successfully

		common.RNRespond(
			w,
			fmt.Sprintf("Successfully uploaded and processed media as %s", mediaFilepath),
			nil,
			http.StatusOK)
	}
}
