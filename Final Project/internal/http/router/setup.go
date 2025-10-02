package router

import (
	"final-project/internal/http/handlers"
)

var repo *handlers.Repository

func New(r *handlers.Repository) {
	repo = r
}
