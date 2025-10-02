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

	mux.HandleFunc("GET /activate", mw.With(repo.ActivateAccount))

	mux.HandleFunc("GET /member/plans", mw.With(repo.ChooseSubscription, mw.Authenticate))
	mux.HandleFunc("GET /member/subscribe", mw.With(repo.Subscribe, mw.Authenticate))

	return mw.Apply(mux)
}
