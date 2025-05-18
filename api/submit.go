package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/yashraj-n/ushort/services"
)

type SubmitBody struct {
	Link string `json:"link"`
}

func HandleSubmit(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var body SubmitBody
	err := decoder.Decode(&body)

	if err != nil {
		http.Error(w, "Failed to parse body", 500)
		return
	}

	if !regexp.MustCompile(`(?i)^(https?:\/\/|www\.)[^\s/$.?#].[^\s]*$`).MatchString(body.Link) {
		http.Error(w, "Provided link does not match regex", http.StatusBadRequest)
		return
	}

	ipAddr := r.Header.Get("X-Forwarded-For")

	if ipAddr == "" {
		// fallback
		ipAddr = "fallback.ushort.internal"
	}

	hash, err := services.CreateRedirectURL(r.Context(), body.Link, ipAddr)

	if err != nil {
		http.Error(w, "Failed to Create Redirect URL", 500)
		return
	}
	fmt.Fprint(w, hash)
}
