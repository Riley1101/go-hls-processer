package api

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"
	config "vid/config"
	src "vid/src"
	ffmpeg "vid/utils/ffmpeg"
	utils "vid/utils/jobqueue"

	"github.com/gorilla/mux"
)

func UploadRoutes(r *mux.Router, db *sql.DB, pool *utils.Pool) {
	r.HandleFunc("/upload", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleUpload(w, r, pool)
	})).Methods("POST")

	r.HandleFunc("/upload", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleGet(w, r, pool)
	})).Methods("GET")
}

func alphabets() error {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	return nil
}

func doWork(pool *utils.Pool) {
	pool.Start()
}

func handleGet(w http.ResponseWriter, r *http.Request, q *utils.Pool) {
	fmt.Fprintf(w, "Job added to queue")

}

func handleUpload(w http.ResponseWriter, r *http.Request, q *utils.Pool) {
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
	file.Upload(multipart_file)
	ffmpeg.CreateHLS(&file, 10)
	fmt.Fprintf(w, "Successfully Uploaded File at %s\n", file.Path)

}
