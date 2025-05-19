package api

import "net/http"

func HandleStaticFiles(w http.ResponseWriter, r *http.Request) {
	fs := http.FileServer(http.Dir("./static"))
	http.StripPrefix("/static/", fs).ServeHTTP(w, r)
}
