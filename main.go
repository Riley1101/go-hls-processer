package main

import (
	"log/slog"
	"net/http"
	api "vid/api"
	config "vid/config"

	"github.com/gorilla/mux"
)

func main() {
	config.InitLogger()
	slog.Info("Server started", "port", 5173)
	db := config.DB{}
	con, _ := db.Connect()
	r := mux.NewRouter()
	api.UploadRoutes(r, con)
	http.ListenAndServe(":5173", r)
}
