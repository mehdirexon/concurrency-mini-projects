package handlers

import (
	"final-project/internal/helpers"
	"final-project/internal/http/render"
	"final-project/internal/models"
	"final-project/internal/service/signer"
	"final-project/internal/shared"
	"final-project/internal/store"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func (a *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, shared.Register, nil)
}

func (a *Repository) PostRegister(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest)
		return
	}

	// TODO Validate Data

	// create a user
	u := store.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Plan:      nil,
	}
	_, err = u.Insert(u)

	if err != nil {
		a.App.Session.Put(r.Context(), shared.Error, true)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// send a activation mail
	url := fmt.Sprintf("http://localhost:8080/activate?email=%s", u.Email)
	signedURL := signer.GenerateTokenFromString(url)
	a.App.InfoLogger.Printf("Signed URL: %s", signedURL)

	msg := models.Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Data:     template.HTML(signedURL),
		Template: "confirmation-email",
	}
	helpers.SendEmail(msg)

	a.App.Session.Put(r.Context(), shared.Flash, "Confirmation email sent.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
