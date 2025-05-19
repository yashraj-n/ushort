package api

import (
	"embed"
	"net/http"
)

//go:embed static/*
var static embed.FS

func HandleStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.FS(static))
	fs.ServeHTTP(w, r)
}
