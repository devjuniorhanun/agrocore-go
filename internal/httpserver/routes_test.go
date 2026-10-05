package httpserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutesHealth(t *testing.T) {
	handler := routes()

	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status code %d, got %d",
			http.StatusOK,
			response.StatusCode,
		)
	}

	var body healthResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}

	if body.Status != "ok" {
		t.Errorf("expected status %q, got %q", "ok", body.Status)
	}
}

func TestRoutesHealthRejectsUnsupportedMethod(t *testing.T) {
	handler := routes()

	request := httptest.NewRequest(http.MethodPost, "/health", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(
			"expected status code %d, got %d",
			http.StatusMethodNotAllowed,
			response.StatusCode,
		)
	}
}

func TestRoutesNotFound(t *testing.T) {
	handler := routes()

	request := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, request)

	response := recorder.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Errorf(
			"expected status code %d, got %d",
			http.StatusNotFound,
			response.StatusCode,
		)
	}
}
