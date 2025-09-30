package router

import (
	"final-project/internal/middlewares"
	"net/http"
)

func Register() http.Handler {

	mux := http.NewServeMux()

	middlewares.Use(middlewares.LoadSession)

	mux.HandleFunc("GET /", middlewares.With(repo.Home, middlewares.LoadSession))

	return mux
}
