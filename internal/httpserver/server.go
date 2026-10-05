package httpserver

import (
	"net/http"
	"time"
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 10 * time.Second
	defaultWriteTimeout      = 10 * time.Second
	defaultIdleTimeout       = 60 * time.Second
)

// New creates an HTTP server configured with the AgroCore routes and
// server-level timeouts.
func New(addr string) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           routes(),
		ReadHeaderTimeout: defaultReadHeaderTimeout,
		ReadTimeout:       defaultReadTimeout,
		WriteTimeout:      defaultWriteTimeout,
		IdleTimeout:       defaultIdleTimeout,
	}
}
