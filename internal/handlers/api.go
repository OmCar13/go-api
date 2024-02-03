package handlers

import (
	"github.com/OmCar13/go-api/internal/middleware"
    "github.com/go-chi/chi"
    chimiddleware "github.com/go-chi/chi/middleware"

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.Use(middleware.Authoriation)

		router.Get("/coins", GetCoinBalance)
	})
}
