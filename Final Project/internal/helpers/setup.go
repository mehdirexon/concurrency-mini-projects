package helpers

import "final-project/internal/config"

var app *config.AppConfig

func Register(a *config.AppConfig) {
	app = a
}
