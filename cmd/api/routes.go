package main

import (
	"expvar"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Use(app.metrics)
	router.Use(app.recoverPanic)
	router.Use(app.enableCORS)
	router.Use(app.rateLimit)
	router.Use(app.authenticate)

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/v1/healthcheck", app.healthcheckHandler)

	router.Route("/v1/movies", func(r chi.Router) {
		r.With(app.requirePermissions("movies:write")).Post("/", app.createMovieHandler)
		r.With(app.requirePermissions("movies:read")).Get("/", app.listMoviesHandler)

		r.With(app.requirePermissions("movies:read")).Get("/{id}", app.showMovieHandler)
		r.With(app.requirePermissions("movies:write")).Patch("/{id}", app.updateMovieHandler)
		r.With(app.requirePermissions("movies:write")).Delete("/{id}", app.deleteMovieHandler)
	})

	router.Route("/v1/users", func(r chi.Router) {
		r.Post("/", app.registerUserHandler)
		r.Put("/activated", app.activateUserHandler)
		r.Put("/password", app.updateUserPasswordHandler)
	})

	router.Route("/v1/tokens", func(r chi.Router) {
		r.Post("/authentication", app.createAuthenticationTokenHandler)
		r.Post("/password-reset", app.createPasswordResetTokenHandler)
	})

	router.Handle("/debug/vars", expvar.Handler())

	return router
}
