package httpserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	addr := ":8080"

	server := New(addr)

	if server.Addr != addr {
		t.Errorf("expected address %q, got %q", addr, server.Addr)
	}

	if server.Handler == nil {
		t.Error("expected server handler to be configured")
	}

	if server.ReadHeaderTimeout != defaultReadHeaderTimeout {
		t.Errorf(
			"expected ReadHeaderTimeout %s, got %s",
			defaultReadHeaderTimeout,
			server.ReadHeaderTimeout,
		)
	}

	if server.ReadTimeout != defaultReadTimeout {
		t.Errorf(
			"expected ReadTimeout %s, got %s",
			defaultReadTimeout,
			server.ReadTimeout,
		)
	}

	if server.WriteTimeout != defaultWriteTimeout {
		t.Errorf(
			"expected WriteTimeout %s, got %s",
			defaultWriteTimeout,
			server.WriteTimeout,
		)
	}

	if server.IdleTimeout != defaultIdleTimeout {
		t.Errorf(
			"expected IdleTimeout %s, got %s",
			defaultIdleTimeout,
			server.IdleTimeout,
		)
	}

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	server.Handler.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf(
			"expected health status code %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}
}
