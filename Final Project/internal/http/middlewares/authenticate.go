package middlewares

import (
	"final-project/internal/helpers"
	"final-project/internal/shared"
	"final-project/internal/store"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsAuthenticated(r) {
			app.Session.Put(r.Context(), shared.Warning, "You are not authenticated")
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}
		user, ok := app.Session.Get(r.Context(), "user").(*store.User)
		if ok {
			print(user)
		}
		next.ServeHTTP(w, r)
	})
}
