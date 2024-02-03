package handlers

import (
	"../internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middlerware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.use(middleware.Authoriation)

		router.Get("/coins", GetCoinBalance)
	})
}
