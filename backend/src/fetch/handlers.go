package fetch

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"rnieb/common"
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
