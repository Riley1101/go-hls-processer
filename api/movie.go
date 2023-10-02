package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	config "vid/config"
	jobqueue "vid/utils/jobqueue"
)

func createMovieTable(db *sql.DB) (sql.Result, error) {
	query := `
    CREATE TABLE movies(
        id INT AUTO_INCREMENT,
        title TEXT NOT NULL,
        file TEXT NOT NULL,
        description TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	res, err := db.Exec(query)
	return res, err
}

func MovieRoutes(r *mux.Router, db *sql.DB, pool *jobqueue.Pool) {

	r.HandleFunc("/api/movie", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		handleGetMovie(w, r, db)
	})).Methods("GET")

	r.HandleFunc("/api/movie/create_table", config.WithLogMiddleware(func(w http.ResponseWriter, r *http.Request) {
		res, err := createMovieTable(db)
		if err != nil {
			json.NewEncoder(w).Encode(err)
		}
		json.NewEncoder(w).Encode(res)
	})).Methods("GET")

	r.HandleFunc("/api/movie", config.WithLogMiddleware(handlePostMovie)).Methods("POST")
}

func handleGetMovie(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Get Movie"))
}
func handlePostMovie(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("POST Movie"))
}
