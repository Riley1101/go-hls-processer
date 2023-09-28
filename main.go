package main

import (
	"github.com/gorilla/mux"
	"net/http"
	api "vid/api"
	config "vid/config"
)

func main() {
	logger := config.InitLogger()
	logger.Info(
		"Server Listening on port 5173",
	)
	r := mux.NewRouter()
	db := config.DB{}
	con, _ := db.Connect()
	api.UploadRoutes(r, con)
	http.ListenAndServe(":5173", nil)

}
