package api

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed index.html
var page string

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, page)
}
