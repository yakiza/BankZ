package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func MainHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/v1/account", NewAccountHandler())

	return r
}
