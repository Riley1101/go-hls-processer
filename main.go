package main

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	api "vid/api"
	config "vid/config"
	jobqueue "vid/utils/jobqueue"
)

func main() {
	config.InitLogger()
	slog.Info("Server started", "port", 5173)
	db := config.DB{}
	con, _ := db.Connect()
	r := mux.NewRouter()
	MAX_WORKERS := 5
	queue := jobqueue.NewQueue("main", MAX_WORKERS)
	defaultWorker := jobqueue.NewWorker(queue)
	pool := jobqueue.NewPool(MAX_WORKERS)

	api.UploadRoutes(r, con, pool)
	http.ListenAndServe(":5173", r)
}
