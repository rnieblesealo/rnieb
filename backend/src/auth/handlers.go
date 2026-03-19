package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"rnieb/common"
)

const (
	JWT_SECRET = "chiikawa"
)

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		// Get submitted credentials

		username := req.FormValue("username")
		pass := req.FormValue("password")

		// Check the password matches

		var actualPass string
		err := db.QueryRow(`
			select password 
			from admins 
			where username = ?
		`, username).Scan(&actualPass)
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Credentials query failed: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		if pass != actualPass {
			common.RNRespond(w, "Invalid credentials", nil, http.StatusUnauthorized)
			return
		}

		common.RNRespond(w, "Password OK!", nil, http.StatusOK)
	}
}
