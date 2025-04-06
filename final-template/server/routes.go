package server

import (
	"net/http"

	appMiddleware "server-http/middleware"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// SetupRoutes initializes and returns the router
func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	// Custom middleware using Zap logger
	r.Use(appMiddleware.LogMiddleware)

	// Add your routes
	r.Get("/", homeHandler)
	r.Get("/about", aboutHandler)

	return r
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	logger := appMiddleware.GetLogger(r.Context())
	logger.Info("Home route accessed")

	w.Write([]byte("Welcome to the home page!"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About page"))
}

// Custom logging middleware using Zap
func loggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("HTTP Request", zap.String("method", r.Method), zap.String("url", r.URL.Path))
			next.ServeHTTP(w, r)
		})
	}
}
