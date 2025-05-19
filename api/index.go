package api

import (
	"io"
	"log/slog"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := http.FS(static).Open("static/index.html")
	if err != nil {
		slog.Error(err.Error())
		http.Error(w, "Failed to read index.html", http.StatusInternalServerError)
		return
	}
	defer htmlFile.Close()
	io.Copy(w, htmlFile)
}
