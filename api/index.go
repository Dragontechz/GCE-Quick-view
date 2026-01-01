package handler

import (
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir("templates"))

	path := filepath.Clean(r.URL.Path)
	if path == "/" {
		path = "/index.html"
	}

	http.StripPrefix("/", fileServer).ServeHTTP(w, r.WithContext(r.Context()))
}
