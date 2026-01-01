package handler

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

var templates embed.FS

func Handler(w http.ResponseWriter, r *http.Request) {
	// Create a sub-filesystem starting at templates
	staticFS, err := fs.Sub(templates, "../templates")
	if err != nil {
		log.Fatal(err)
	}

	// Serve static files (including index.html for directories)
	fileServer := http.FileServer(http.FS(staticFS))

	// Clean the path to prevent directory traversal
	path := filepath.Clean(r.URL.Path)

	// Optional: Serve index.html for clean URLs (e.g., /about/ → /about.html)
	if path == "/" {
		path = "/index.html"
	} else if info, err := fs.Stat(staticFS, path); err == nil && info.IsDir() {
		path = filepath.Join(path, "index.html")
	}

	// Serve the file
	http.StripPrefix("/", fileServer).ServeHTTP(w, r.WithContext(r.Context()))
}
