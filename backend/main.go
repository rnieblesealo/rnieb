package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"rnieb/auth"
	"rnieb/common"
	"rnieb/fetch"
	"rnieb/upload"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gographics/imagick.v3/imagick"
)

const (
	PORT = ":8080"
)

// NOTE: http.handlerfunc implements http.handler intf

// Allows any origin to access this; effectively we're a public API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow any origin

		w.Header().Set("Access-Control-Allow-Origin", "*")

		// allow get, post delete only

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE")

		// allow authorization and content type headers

		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		// TODO: not sure what this does/is

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get the encrypted jwt string attached by the client

		jwtString := r.Header.Get("Authorization")
		if jwtString == "" {
			common.RNRespond(
				w,
				"Invalid auth token",
				nil,
				http.StatusUnauthorized)

			return
		}

		// decrypt token ( do signature verification with our secret )

		jwt, err := jwt.Parse(
			jwtString,
			func(token *jwt.Token) (interface{}, error) {

				// NOTE: this is a callback that returns the jwt secret to use in decrypting

				/* since we may have more than 1 secret, etc., this allows flexibility
				...we can inspect the token's header first and then select the right secret!
				( as opposed to just passing one key directly ) */

				// i'm just passing the only secret i have bc my case is simple tho

				return []byte(os.Getenv("JWT_SECRET")), nil
				// WARNING: not bytecasting this will give SNEAKY errors
			})

		if err != nil {
			common.RNRespond(
				w,
				fmt.Sprintf("JWT parsing failed: %s", err),
				nil,
				http.StatusBadRequest)

			return
		}

		// check that it is valid ( we were able to decrypt it )

		if !jwt.Valid {
			common.RNRespond(
				w,
				"JWT not valid",
				nil,
				http.StatusBadRequest)

			return
		}

		// run next middleware

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize ImageMagick
	/* Must be done once; C API */

	imagick.Initialize()
	defer imagick.Terminate()

	// Get important env variables

	dbPath := os.Getenv("DB_PATH")
	uploadPath := os.Getenv("UPLOAD_PATH")
	secretsPath := os.Getenv("DOTENV_PATH")

	// Load secrets .env
	// Non-secret environment variables are loaded by compose

	godotenv.Load(secretsPath)

	fmt.Printf("%s %s %s\n", dbPath, uploadPath, secretsPath)

	// Open database connection

	db, err := sql.Open("sqlite3", dbPath) // Open DB connection
	if err != nil {
		log.Fatalf("Failed to open DB connection: %s\n", err)
	}
	defer db.Close()

	// Setup auth handlers

	http.HandleFunc("/api/login", auth.Login(db))
	http.Handle("/api/me", authMiddleware(upload.Ping("You are logged in!"))) // Login check

	// Setup upload handlers (auth-protected)

	http.Handle("/api/ping", upload.Ping("Marco? Polo!"))
	http.Handle("/api/upload", authMiddleware(upload.UploadImage(db, uploadPath)))
	http.Handle("/api/upload-video", upload.UploadVideo(db, uploadPath))

	// Setup fetch handlers

	http.HandleFunc("/api/get-drawings", fetch.GetDrawings(db))
	http.Handle("/api/delete-drawing", authMiddleware(fetch.DeleteDrawing(db, uploadPath)))

	// Setup image fileserver

	http.Handle("/api/uploads/", // Setup handler for uploads route
		http.StripPrefix("/api/uploads/", // Strip this prefix from URL ( Leaves only filename )
			http.FileServer(http.Dir(uploadPath)))) // Look for that file in /uploads

	fmt.Printf("Starting rnieb server on port %s...\n", PORT)

	// Start HTTP

	log.Fatal(http.ListenAndServe(PORT,
		corsMiddleware(
			http.DefaultServeMux)))

	// NOTE: ListenAndServe always returns non nil err; only fails if it errors???
}
