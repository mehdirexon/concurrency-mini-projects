package router

import (
	"final-project/internal/handlers"
)

var repo *handlers.Repository

func New(r *handlers.Repository) {
	repo = r
}
