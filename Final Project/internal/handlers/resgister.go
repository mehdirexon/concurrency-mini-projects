package handlers

import (
	"final-project/internal/render"
	"final-project/internal/shared"
	"net/http"
)

func (a *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, shared.Register, nil)
}

func (a *Repository) PostRegister(w http.ResponseWriter, r *http.Request) {
<<<<<<< Updated upstream

=======
	// create a user

	// send a activation mail

	// subscribe the user to an account
>>>>>>> Stashed changes
}
