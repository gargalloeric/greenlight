package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	shutdownErrors := make(chan error)

	// Start goroutine to listen for termination signals
	go func() {
		// Create the quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)

		// Use signal.Notify to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel. Any other signals will not be caught by
		// signal.Notify and will retain their default behavior.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// This code blocks until it reads the signal from the channel.
		s := <-quit

		app.logger.Info("shutting down server", "signal", s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Call Shutdown() on our server, passing in the context we just made.
		// Shutdown() will return nil if the graceful shutdown was successful, or an
		// error (which may happen because of a problem closing the listeners, or
		// because the shutdown didn't complete before the 30-second context deadline is
		// hit).
		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownErrors <- err
		}

		app.logger.Info("completing background tasks", "addr", srv.Addr)

		app.wg.Wait()
		shutdownErrors <- nil
	}()

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	if err := <-shutdownErrors; err != nil {
		return err
	}

	app.logger.Info("server stopped", "addr", srv.Addr)

	return nil
}
