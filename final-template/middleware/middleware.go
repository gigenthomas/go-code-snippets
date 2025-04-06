package middleware

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

// LoggerContextKey defines a key for the logger in the context.
type LoggerContextKey string

// LogMiddleware is a middleware that injects the logger into the request context.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the logger from the global logger or create a new one for this request
		logger := zap.L().With(zap.String("method", r.Method), zap.String("path", r.URL.Path))

		// Inject the logger into the request context
		ctx := context.WithValue(r.Context(), LoggerContextKey("logger"), logger)

		// Pass the request with the new context to the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetLogger retrieves the logger from the request context
func GetLogger(ctx context.Context) *zap.Logger {
	if logger, ok := ctx.Value(LoggerContextKey("logger")).(*zap.Logger); ok {
		return logger
	}
	// Return a default logger if not found
	return zap.L()
}
