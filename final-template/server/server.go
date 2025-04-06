package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type HTTPServer struct {
	server *http.Server
}

type Server interface {
	Start() error
	Stop() error
}

func NewHTTPServer(addr string, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		server: &http.Server{
			Addr:    addr,
			Handler: handler,
		},
	}
}

func (h *HTTPServer) Start() error {
	fmt.Println("Starting server on", h.server.Addr)
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()
	return nil
}

func (h *HTTPServer) Stop() error {
	fmt.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return h.server.Shutdown(ctx)
}
