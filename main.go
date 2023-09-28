package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	api "vid/api"
	config "vid/config"
)

func main() {
	fmt.Println("Server started")
	config.InitLogger()
	r := mux.NewRouter()
	db := config.DB{}
	con, _ := db.Connect()
	api.UploadRoutes(r, con)
	http.ListenAndServe(":5173", nil)

}
