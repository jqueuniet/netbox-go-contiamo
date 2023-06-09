package router

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// NewRouter creates a new router for the spec and the given handlers.
// NetBox REST API
//
// 3.5.1 (3.5)
func NewRouter() http.Handler {

	r := chi.NewRouter()

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
