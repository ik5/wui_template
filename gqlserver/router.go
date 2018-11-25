package gqlserver

import (
	"net/http"

	"github.com/go-chi/chi"
)

func auth(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

// GQLRouter is creating graphql router
func GQLRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(auth)
	r.Post("/", root)
	return r
}
