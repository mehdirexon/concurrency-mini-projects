package handlers

import (
	"final-project/internal/helpers"
	"final-project/internal/http/render"
	"final-project/internal/models"
	"final-project/internal/shared"
	"net/http"
)

func (a *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, shared.Login, nil)
}

func (a *Repository) PostLogin(w http.ResponseWriter, r *http.Request) {
	_ = a.App.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := a.Store.GetUserByEmail(email)
	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	validPassword, err := a.Store.UserPasswordMatches(user, password)
	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		msg := models.Message{
			To:      email,
			Subject: "failed log in attempt!",
			Data:    "Invalid login attempt!",
		}

		helpers.SendEmail(msg)

		a.App.Session.Put(r.Context(), shared.Error, "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if user.Active != 1 {
		a.App.Session.Put(r.Context(), shared.Error, "Account is not activated. check your email.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	a.App.Session.Put(r.Context(), "userID", user.ID)
	a.App.Session.Put(r.Context(), "user", user)

	a.App.Session.Put(r.Context(), shared.Flash, "Successful logged!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (a *Repository) Logout(w http.ResponseWriter, r *http.Request) {
	_ = a.App.Session.Destroy(r.Context())
	_ = a.App.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
