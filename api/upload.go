package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func UploadRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/upload", handleUpload).Methods("POST")
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload")
}
