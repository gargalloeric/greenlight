package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(app.authenticate)
	router.Use(app.rateLimit)
	router.Use(app.recoverPanic)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Route("/v1/movies", func(r chi.Router) {
		r.Post("/", app.createMovieHandler)
		r.Get("/", app.listMoviesHandler)

		r.Get("/{id}", app.showMovieHandler)
		r.Patch("/{id}", app.updateMovieHandler)
		r.Delete("/{id}", app.deleteMovieHandler)
	})

	router.Route("/v1/users", func(r chi.Router) {
		r.Post("/", app.registerUserHandler)
		r.Put("/activated", app.activateUserHandler)
	})

	router.Route("/v1/tokens", func(r chi.Router) {
		r.Post("/authentication", app.createAuthenticationTokenHandler)
	})

	return router
}
