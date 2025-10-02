package handlers

import (
	"final-project/internal/shared"
	"final-project/internal/signer"
	"fmt"
	"net/http"
)

func (a *Repository) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// validate url
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost:8080%s", url)
	okay := signer.VerifyToken(testURL)

	if !okay {
		a.App.Session.Put(r.Context(), shared.Error, "Invalid token")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	user, err := a.Store.GetUserByEmail(r.URL.Query().Get("email"))
	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, "user not found!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	user.Active = 1
	err = user.Update()
	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, "unable to activate")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	a.App.Session.Put(r.Context(), shared.Flash, "Account activated!")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	// generate an invoice

	// send an email with attachments

	// send an email with the invoice attached

	// subscribe the user to an account

}
