package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func ClientError(w http.ResponseWriter, status int) {
	app.ErrorLogger.Printf("Client error: %d", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLogger.Printf("Server error: %d", trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
