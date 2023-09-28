package config

import (
	"log/slog"
	"net/http"
)

func ServeStatic() {
	slog.Info("Serving static files", "Path", "/assets/")

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
