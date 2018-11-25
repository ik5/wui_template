package restserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ik5/wui_template/auth"
)

func tokenAuth(handle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		token, ok := r.Header["X-Api-Token"]
		if !ok || !auth.FindToken(r.Context(), token[0]) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		handle.ServeHTTP(w, r)
	})
}

// RestRouter register routing for REST requests
func RestRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(tokenAuth)
	r.Route("/v1", func(router chi.Router) {
	})

	return r
}
