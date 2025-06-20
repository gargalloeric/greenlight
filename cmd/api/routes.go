package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(app.recoverPanic)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Route("/v1/movies", func(r chi.Router) {
		r.Post("/", app.createMovieHandler)
		r.Get("/{id}", app.showMovieHandler)
		r.Patch("/{id}", app.updateMovieHandler)
		r.Delete("/{id}", app.deleteMovieHandler)
	})

	return router
}
