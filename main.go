package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/devjuniorhanun/agrocore-go/internal/config"
	"github.com/devjuniorhanun/agrocore-go/internal/httpserver"
)

const shutdownTimeout = 10 * time.Second

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	addr := net.JoinHostPort(
		cfg.HTTP.Host,
		strconv.Itoa(cfg.HTTP.Port),
	)

	server := httpserver.New(addr)

	shutdownSignal, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	serverError := make(chan error, 1)

	go func() {
		log.Printf("starting AgroCore HTTP server on %s", server.Addr)

		serverError <- server.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server failed: %v", err)
		}

	case <-shutdownSignal.Done():
		log.Println("shutdown signal received")
	}

	shutdownContext, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeout,
	)
	defer cancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Printf("HTTP server shutdown failed: %v", err)
	}

	log.Println("AgroCore HTTP server stopped")
}
