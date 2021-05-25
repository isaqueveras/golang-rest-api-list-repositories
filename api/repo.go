package handlers

import (
	"net/http"
)

func RepoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Oii estou aqui ;)"))
}
