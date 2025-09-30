package helpers

import "net/http"

var IsAuthenticated = func(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}
