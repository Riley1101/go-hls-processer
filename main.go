package main

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	api "vid/api"
	config "vid/config"
	jobqueue "vid/utils/jobqueue"
	web "vid/website"
)

func main() {
	config.InitLogger()
	slog.Info("Server started", "port", 5173)
	db := config.DB{}
	con, _ := db.Connect()
	r := mux.NewRouter()
	MAX_WORKERS := 5
	pool := jobqueue.NewPool(MAX_WORKERS)
	web.WebRoutes(r)
	api.UploadRoutes(r, con, pool)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./processed/")))
	http.ListenAndServe(":5173", r)
}

func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
