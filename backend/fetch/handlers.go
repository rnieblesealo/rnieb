package fetch

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"rnieb/common"
	"slices"

	_ "github.com/mattn/go-sqlite3"
)

type Media struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

// Gets all media of the specified type
func GetAllMediaOfType(db *sql.DB, mediaType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Ensure passed type is allowed!

		if !slices.Contains(common.GetAllowedDBMediaTypes(), mediaType) {
			common.RNRespond(
				w,
				fmt.Sprintf("Illegal media type '%s'", mediaType),
				nil,
				http.StatusBadRequest,
			)

			return
		}

		// Query for all media of that type

		rows, err := db.Query(`
    	SELECT id, name, description, filename FROM media WHERE type = ?
		`, mediaType)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Error getting media: %s", err),
				nil,
				http.StatusInternalServerError,
			)

			return
		}
		defer rows.Close() // Query results are kept open; we must close them

		// Marshal list of media

		var media []Media

		for rows.Next() {
			var mediaItem Media

			rows.Scan(
				&mediaItem.ID,
				&mediaItem.Name,
				&mediaItem.Description,
				&mediaItem.Path,
			) // Values are scanned into Go with closest type to the DB's

			media = append(media, mediaItem)
		}

		common.RNRespond(
			w,
			fmt.Sprintf("Successfully retrieved all media of type '%s'", mediaType),
			media,
			http.StatusOK,
		)
	}
}

// Deletes media by ID from the upload path
func DeleteMedia(db *sql.DB, uploadPath string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get the deletion ID from URL params

		deletionID := req.URL.Query().Get("id") // Get deletion ID query param

		// Get the path of the image we want to delete

		var deletionImageFilename string
		db.QueryRow(`
		SELECT filename 
		FROM media 
		WHERE id = ?
		`, deletionID).Scan(&deletionImageFilename) // NOTE: Only db.Query requires closing

		// Delete the image from the filesystem

		err := os.Remove(filepath.Join(uploadPath, deletionImageFilename))

		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to delete image (location %s): %s", deletionImageFilename, err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Delete the record from the DB

		_, err = db.Exec(`
		DELETE FROM media WHERE id = ?	
	`, deletionID)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to delete record (ID %s): %s", deletionID, err),
				nil,
				http.StatusInternalServerError)

			return
		}

		// Respond successfully

		common.RNRespond(
			w,
			fmt.Sprintf(
				"Successfully deleted ID %s ( file: %s )",
				deletionID,
				deletionImageFilename),
			nil,
			http.StatusOK,
		)
	}
}
