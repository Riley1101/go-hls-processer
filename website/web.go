package website

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
)

func WebRoutes(r *mux.Router) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, tmpl)
	})

}

func HomeHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	processedDir, err := os.ReadDir("processed")
	if err != nil {
		panic(err)
	}
	var files []string
	for _, file := range processedDir {
		files = append(files, file.Name())
	}

	// send the processed files to the template
	t.Execute(w, files)
}
