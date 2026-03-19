package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"rnieb/common"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	JWT_SECRET = "chiikawa" // WARNING: temporary!
)

func Login(db *sql.DB) http.HandlerFunc {
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
				http.StatusBadRequest)

			return
		}

		if pass != actualPass {
			common.RNRespond(w, "Invalid credentials", nil, http.StatusUnauthorized)
			return
		}

		// Return auth token
		/* given back as json under data: {"token":<tokenstring>} */

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(24 * time.Hour).Unix(), // valid for 24 hours
		})

		tokenString, err := token.SignedString([]byte(JWT_SECRET)) // Sign token with JWT secret
		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("Failed to sign token: %s", err),
				nil,
				http.StatusInternalServerError)

			return
		}

		type TokenResponse struct {
			Token string `json:"token"`
		}

		common.RNRespond(
			w,
			"Login successful",
			TokenResponse{Token: tokenString},
			http.StatusOK)
	}
}
