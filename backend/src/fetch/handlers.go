package fetch

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"rnieb/common"

	_ "github.com/mattn/go-sqlite3"
)

type Drawing struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Path        string `json:"path"`
}

func GetDrawings(w http.ResponseWriter, r *http.Request) {
	// Connect to DB

	db, err := sql.Open("sqlite3", "./rnieb.db")
	if err != nil {
		common.RNRespond(
			w,
			fmt.Sprintf("Failed to connect to DB: %s", err),
			nil,
			http.StatusInternalServerError,
		)

		return
	}
	defer db.Close()

	// Query for all rows of drawings table
	/* Only include the id, name, description and path, no need for creation date */

	rows, err := db.Query(`
			SELECT id, name, description, path FROM drawings
	`)
	if err != nil {
		common.RNRespond(
			w,
			fmt.Sprintf("Drawings query error: %s", err),
			nil,
			http.StatusInternalServerError,
		)

		return
	}
	defer rows.Close() // Query results are kept open; we must close them

	// Marshal list of drawings using a go type

	var drawings []Drawing

	for rows.Next() {
		var drawing Drawing

		rows.Scan(
			&drawing.ID,
			&drawing.Name,
			&drawing.Description,
			&drawing.Path,
		) // Values are scanned into Go with closest type to the DB's

		drawings = append(drawings, drawing)
	}

	common.RNRespond(
		w,
		"Successfully retrieved drawings",
		drawings,
		http.StatusOK,
	)
}

// Deletes a drawing and its file
func DeleteDrawing(w http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("sqlite3", "./rnieb.db")
	if err != nil {
		common.RNRespond(
			w,
			fmt.Sprintf("Failed to connect to DB: %s", err),
			nil,
			http.StatusInternalServerError)

		return
	}
	defer db.Close()

	// Get the deletion ID from URL params

	deletionID := req.URL.Query().Get("id") // Get deletion ID query param

	// Get the path of the image we want to delete

	var deletionImagePath string
	db.QueryRow(`
		SELECT path
		FROM drawings
		WHERE id = ?
		`, deletionID).Scan(&deletionImagePath) // NOTE: Only db.Query requires closing

	// Delete the image from the filesystem

	err = os.Remove(deletionImagePath)
	if err != nil {
		common.RNRespond(
			w,
			fmt.Sprintf("Failed to delete image (location %s): %s", deletionImagePath, err),
			nil,
			http.StatusInternalServerError)

		return
	}

	// Delete the record from the DB

	_, err = db.Exec(`
		DELETE FROM drawings WHERE id = ?	
	`, deletionID)
	if err != nil {
		common.RNRespond(
			w,
			fmt.Sprintf("Failed to delete record (ID %s): %s", deletionID, err),
			nil,
			http.StatusInternalServerError)

		return
	}

	common.RNRespond(
		w,
		fmt.Sprintf(
			"Successfully deleted ID %s with image %s",
			deletionID,
			deletionImagePath),
		nil,
		http.StatusOK,
	)
}
