package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	config "vid/config"
	src "vid/src"

	"github.com/gorilla/mux"
)

func UploadRoutes(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/upload", config.WithLogMiddleware(handleUpload)).Methods("POST")
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	multipart_file, file_headers, err := r.FormFile("file")
	file := src.File{
		OriginalName: file_headers.Filename,
		Size:         file_headers.Size,
		Header:       file_headers.Header.Get("Content-Type"),
	}
	err = src.ValidateDir("uploads")
	file.Upload(multipart_file)
	if err != nil {
		slog.Error("Error Retrieving the File")
		return
	}
	defer multipart_file.Close()
	fmt.Fprintf(w, "File Uploaded Successfully")

}
