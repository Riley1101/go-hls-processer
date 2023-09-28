package website

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func WebRoutes(r *mux.Router) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, tmpl)
	})

}

func HomeHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(w, nil)
}
