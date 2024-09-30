package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/moabdelazem/sample-app/internal/store"
)

// Application is a struct that contains all the necessary components for the application
type Application struct {
	Config  Config
	Storage store.Storage
}

// Config is a struct that contains the configuration for the application
type Config struct {
	Address string
	DB      DBConfig
}

// DBConfig is a struct that contains the configuration for the database
type DBConfig struct {
	Address             string
	MaxOpenConnections  int
	MaxIdleConnnections int
	MaxIdleTime         time.Duration
}

// Mount is a method on the Application struct that returns a new http.ServeMux instance
//   - This method is responsible for defining the routes for the application
//   - The http.ServeMux instance is a multiplexer that matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL
func (app *Application) Mount() *chi.Mux {

	// Create a new http.ServeMux instance
	router := chi.NewRouter()

	// Enable the middlewares for the router
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	// Define New Group for v1 API
	router.Route("/v1/api", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Welcome to the API"))
		})
		// Define a route for the health URL path
		r.Get("/health", app.HealthCheckHandler)

		r.Route("/users", func(r chi.Router) {
			// Define a route for the users URL path
			r.Get("/", app.GetUsersHandler)
			r.Get("/{id}", app.GetUserByIDHandler)
			r.Get("/{username}", app.GetUserByUsernameHandler)

			// Post Request
			r.Post("/", app.UserRegisterHandler)
		})

	})

	return router
}

// Run is a method on the Application struct that starts the HTTP server
// - This method is responsible for starting the HTTP server and listening for incoming requests
// - The http.Server instance is created with the Addr, Handler, WriteTimeout, ReadTimeout, and IdleTimeout fields set
func (app *Application) Run(router *chi.Mux) error {

	// Create a new http.Server instance
	srv := &http.Server{
		Addr:         app.Config.Address,
		Handler:      router,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	// Log a message to indicate that the server is listening
	log.Printf("Server is listening on %s", app.Config.Address)

	// Start the HTTP server
	return srv.ListenAndServe()
}
