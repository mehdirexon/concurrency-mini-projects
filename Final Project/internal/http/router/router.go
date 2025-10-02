package router

import (
	"final-project/internal/http/middlewares"
	"net/http"
)

func Register() http.Handler {

	mux := http.NewServeMux()

	middlewares.Use(middlewares.LoadSession)

	mux.HandleFunc("GET /", middlewares.With(repo.Home))

	mux.HandleFunc("GET /login", middlewares.With(repo.Login))
	mux.HandleFunc("POST /login", middlewares.With(repo.PostLogin))
	mux.HandleFunc("GET /logout", middlewares.With(repo.Logout))

	mux.HandleFunc("GET /register", middlewares.With(repo.Register))
	mux.HandleFunc("POST /register", middlewares.With(repo.PostRegister))

	mux.HandleFunc("GET /activate", middlewares.With(repo.ActivateAccount))

	mux.HandleFunc("GET /member/plans", middlewares.With(repo.ChooseSubscription, middlewares.Authenticate))
	mux.HandleFunc("GET /member/subscribe", middlewares.With(repo.Subscribe, middlewares.Authenticate))

	return middlewares.Apply(mux)
}
