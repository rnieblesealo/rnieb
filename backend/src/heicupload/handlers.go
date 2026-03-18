package heicupload

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	UPLOAD_FIELD_NAME = "image"
	UPLOAD_DIR        = "/uploads"
	MAX_FORM_SIZE     = 32000000 // 32 mb upload limit; parsemultipartform takes bytes
)

// Used to check for heartbeat
func PingHandler(w http.ResponseWriter, req *http.Request) {
	resp := map[string]interface{}{ // [keyType]valueType; empty interface = any type
		"messageType": "S",
		"message":     "",
		"data":        "PONG",
	}

	w.Header().Set("Content-Type", "application/json") // set the headers
	w.WriteHeader(http.StatusOK)                       // write the set headers and attach a statuscode

	json.NewEncoder(w).Encode(resp) // write json response
}

// Handles image uploads
func UploadHandler(w http.ResponseWriter, req *http.Request) {

	// CHECK FOR FORM PARSING ERRORS

	if err := req.ParseMultipartForm(MAX_FORM_SIZE); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// GET REF TO FILE HEADERS;
	/* only accessible after calling ParseMultipartForm
	it seems file headers are handles to an uploaded file */

	fmt.Printf("Hi!")

	imageName := req.FormValue("name")
	imageDescription := req.FormValue("description")
	_, imageFile, err := req.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var errNew string = ""
	var httpStatus int = 0

	// GO OVER ALL FILE HANDLES AND RECEIVE EACH...

	// OPEN THE FILE POINTED TO BY HANDLER

	file, err := imageFile.Open()
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	defer file.Close()

	// READ FIRST 512 BYTES OF THE FILE
	/* the mime type (i.e. file's type) can be identified from the first 512 bytes
	we do that in the step after this! */

	buf := make([]byte, 512)
	_, err = file.Read(buf)
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	// CHECK THE CONTENT TYPE
	/* ( we only want images of heic, png, or jpg type ) */

	fileMimeType := mimetype.Detect(buf).String()

	fmt.Printf("Detected filetype for %s: %s\n", imageFile.Filename, fileMimeType)

	allowedMimeTypes := []string{
		"image/png",
		"image/jpeg",
		"image/heic",
	}

	if !mimetype.EqualsAny(fileMimeType, allowedMimeTypes...) {
		errNew = "Uploaded image must be PNG, JPG, or HEIC"
		httpStatus = http.StatusBadRequest
	}

	// MOVE READ CURSOR TO BEGINNING OF FILE

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	// CREATE UPLOADS DIR
	/* ( modeperm = full read, write + execute perms; same as chmod 777 )
	remember that executing a dir = opening it ) */

	err = os.MkdirAll(UPLOAD_DIR, os.ModePerm)
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	// CREATE THE IMAGE FILE

	imgFile, err := os.Create(fmt.Sprintf("%s/%d%s", UPLOAD_DIR, time.Now().UnixNano(), filepath.Ext(imageFile.Filename)))
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	defer imgFile.Close()

	// COPY UPLOADED IMAGE DATA TO IMAGE FILE

	_, err = io.Copy(imgFile, file)
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	}

	// CONVERT THE IMAGE IF NOT IN PNG FORMAT

	var imageFilepath string

	if fileMimeType != "image/png" {
		convertedFilepath, err := ConvertToPNG(imgFile.Name())
		if err != nil {
			errNew = err.Error()
			httpStatus = http.StatusInternalServerError
		}

		imageFilepath, _ = ResizePNG(convertedFilepath)
	} else {
		// resize only if already png
		imageFilepath, _ = ResizePNG(imgFile.Name())
	}

	// CREATE DB ENTRY FOR IT

	db, err := sql.Open("sqlite3", "./rnieb.db")
	if err != nil {
		errNew = err.Error()
		httpStatus = http.StatusInternalServerError
	} else {
		defer db.Close()
		db.Exec(`
			INSERT INTO drawings (name, description, path)
			VALUES (?, ?, ?)	
	`, imageName, imageDescription, imageFilepath)
	}

	// RESPOND

	message := "File uploaded successfully"
	messageType := "S"

	if errNew != "" {
		message = errNew
		messageType = "E"
	}

	if httpStatus == 0 {
		httpStatus = http.StatusOK
	}

	resp := map[string]interface{}{
		"messageType": messageType,
		"message":     message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	json.NewEncoder(w).Encode(resp)
}
