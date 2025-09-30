package middlewares

import (
	config "final-project/internal/config"
	"net/http"
)

var app *config.AppConfig
var middlewares []func(http.Handler) http.Handler

func Register(a *config.AppConfig) {
	app = a
}
