package router

import (
	mw "final-project/internal/middlewares"
	"net/http"
)

func Register() http.Handler {

	mux := http.NewServeMux()

	mw.Use(mw.LoadSession)

	mux.HandleFunc("GET /", mw.With(repo.Home))

	mux.HandleFunc("GET /login", mw.With(repo.Login))
	mux.HandleFunc("POST /login", mw.With(repo.PostLogin))
	mux.HandleFunc("GET /logout", mw.With(repo.Logout))

	mux.HandleFunc("GET /register", mw.With(repo.Register))
	mux.HandleFunc("POST /register", mw.With(repo.PostRegister))

	mux.HandleFunc("POST /activate-account", mw.With(repo.ActivateAccount))

	return mw.Apply(mux)
}
