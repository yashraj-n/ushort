package api

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yashraj-n/ushort/services"
)

func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	hash := chi.URLParam(r, "hash")
	link, err := services.GetRedirectURL(r.Context(), hash)
	if err != nil {
		http.Error(w, "Something went wrong!", 500)
		slog.Warn("Error while redirecting", "error", err.Error())
		return
	}
	http.Redirect(w, r, link, http.StatusMovedPermanently)
}
