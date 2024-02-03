package handlers

import (
	"github.com/OmCar13/go-api/internal/middleware"
	"github.com/go-chi/chi"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.Use(middleware.Authoriation)

		router.Get("/coins", GetCoinBalance)
	})
}
