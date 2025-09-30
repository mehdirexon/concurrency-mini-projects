package middlewares

import "net/http"

func LoadSession(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
