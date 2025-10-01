package handlers

import (
	"final-project/internal/config"
	"final-project/internal/store"
)

var repo Repository

type Repository struct {
	App   *config.AppConfig
	Store store.Store
}

func Register(a *config.AppConfig, s store.Store) {
	repo.App = a
	repo.Store = s
}

func GetRepo() *Repository {
	return &repo
}
