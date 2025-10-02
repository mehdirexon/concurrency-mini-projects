package handlers

import (
	"final-project/internal/http/render"
	"final-project/internal/shared"
	"net/http"
)

func (a *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, shared.Home, nil)
}
