package config

import "testing"

func TestLoadDefaults(t *testing.T) {
	t.Setenv("HTTP_HOST", "")
	t.Setenv("HTTP_PORT", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected configuration to load, got error: %v", err)
	}

	if cfg.HTTP.Host != defaultHTTPHost {
		t.Errorf(
			"expected default HTTP host %q, got %q",
			defaultHTTPHost,
			cfg.HTTP.Host,
		)
	}

	if cfg.HTTP.Port != defaultHTTPPort {
		t.Errorf(
			"expected default HTTP port %d, got %d",
			defaultHTTPPort,
			cfg.HTTP.Port,
		)
	}
}

func TestLoadFromEnvironment(t *testing.T) {
	t.Setenv("HTTP_HOST", "127.0.0.1")
	t.Setenv("HTTP_PORT", "9000")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected configuration to load, got error: %v", err)
	}

	if cfg.HTTP.Host != "127.0.0.1" {
		t.Errorf(
			"expected HTTP host %q, got %q",
			"127.0.0.1",
			cfg.HTTP.Host,
		)
	}

	if cfg.HTTP.Port != 9000 {
		t.Errorf(
			"expected HTTP port %d, got %d",
			9000,
			cfg.HTTP.Port,
		)
	}
}

func TestLoadRejectsInvalidPort(t *testing.T) {
	tests := []struct {
		name string
		port string
	}{
		{
			name: "not a number",
			port: "invalid",
		},
		{
			name: "zero",
			port: "0",
		},
		{
			name: "negative",
			port: "-1",
		},
		{
			name: "greater than maximum",
			port: "65536",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("HTTP_PORT", tt.port)

			if _, err := Load(); err == nil {
				t.Fatalf(
					"expected HTTP_PORT %q to be rejected",
					tt.port,
				)
			}
		})
	}
}
