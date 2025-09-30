package handlers

import (
	config "final-project/internal/config"
	"final-project/internal/store"
)

type Repository struct {
	App   *config.AppConfig
	Store *store.Store
}
