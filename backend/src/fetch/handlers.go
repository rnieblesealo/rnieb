package fetch

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"os"
)

// TODO: structure the responses similarly to other http

func ListImages(w http.ResponseWriter, r *http.Request) {
	dirEntries, err := os.ReadDir("/uploads")
	if err != nil {
		fmt.Printf("ListImages: %s\n", err.Error())
		return
	}

	var names []string
	for _, dirEntry := range dirEntries {
		names = append(names, dirEntry.Name())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(names)
}

func GetDrawings(w http.ResponseWriter, r *http.Request) {
	// connect to db
	db, err := sql.Open("sqlite3", "./rnieb.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// query for drawings rows
	rows, err := db.Query(`
			SELECT id, name, description, path FROM drawings
	`)
	defer rows.Close() // query results are kept open as a cursor (?); we must close

	// marshal list of drawings using a go type
	type Drawing struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Path        string `json:"path"`
	}

	var drawings []Drawing
	for rows.Next() {
		var drawing Drawing

		rows.Scan(
			&drawing.ID,
			&drawing.Name,
			&drawing.Description,
			&drawing.Path,
		)
		// values are scanned into go type that resembles the db's closest

		drawings = append(drawings, drawing)
	}

	// send json back
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(drawings)
}
