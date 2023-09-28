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
	if err != nil {
		slog.Error("Error Retrieving the File")
		return
	}

	file := src.File{
		OriginalName: file_headers.Filename,
		Size:         file_headers.Size,
		Header:       file_headers.Header.Get("Content-Type"),
	}

	src.ValidateDir("uploads")
	src.ValidateFileType(multipart_file, w)

	file.Upload(multipart_file)
	fmt.Fprintf(w, "Successfully Uploaded File at %s\n", file.Path)
	defer multipart_file.Close()

}
