package agriculturalyear

import (
	"errors"
	"testing"
	"time"

	"github.com/devjuniorhanun/agrocore-go/internal/shared/date"
)

func TestNewAgriculturalYear(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantName string
		wantErr  error
	}{
		{
			name:     "creates agricultural year",
			input:    "2026/2027",
			wantName: "2026/2027",
		},
		{
			name:     "trims agricultural year name",
			input:    "  2026/2027  ",
			wantName: "2026/2027",
		},
		{
			name:    "rejects empty name",
			input:   "",
			wantErr: ErrNameRequired,
		},
		{
			name:    "rejects whitespace-only name",
			input:   "   ",
			wantErr: ErrNameRequired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.input)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("New() error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				return
			}

			if got.ID() != 0 {
				t.Errorf("ID() = %d, want 0", got.ID())
			}

			if got.Name() != tt.wantName {
				t.Errorf("Name() = %q, want %q", got.Name(), tt.wantName)
			}

			if got.Status() != StatusActive {
				t.Errorf(
					"Status() = %q, want %q",
					got.Status(),
					StatusActive,
				)
			}
		})
	}
}

func TestRestoreAgriculturalYear(t *testing.T) {
	got, err := Restore(
		10,
		"2026/2027",
		StatusActive,
	)
	if err != nil {
		t.Fatalf("Restore() unexpected error: %v", err)
	}

	if got.ID() != 10 {
		t.Errorf("ID() = %d, want 10", got.ID())
	}

	if got.Name() != "2026/2027" {
		t.Errorf(
			"Name() = %q, want %q",
			got.Name(),
			"2026/2027",
		)
	}

	if got.Status() != StatusActive {
		t.Errorf(
			"Status() = %q, want %q",
			got.Status(),
			StatusActive,
		)
	}
}

func TestRestoreRejectsEmptyName(t *testing.T) {
	_, err := Restore(
		10,
		"   ",
		StatusActive,
	)

	if !errors.Is(err, ErrNameRequired) {
		t.Fatalf(
			"Restore() error = %v, want %v",
			err,
			ErrNameRequired,
		)
	}
}

func TestAgriculturalYearDates(t *testing.T) {
	agriculturalYear, err := New("2026/2027")
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if _, ok := agriculturalYear.OpeningDate(); ok {
		t.Error("expected opening date to be undefined")
	}

	if _, ok := agriculturalYear.ClosingDate(); ok {
		t.Error("expected closing date to be undefined")
	}

	openingDate, err := date.New(2026, time.July, 1)
	if err != nil {
		t.Fatalf("date.New() unexpected error: %v", err)
	}

	closingDate, err := date.New(2027, time.June, 30)
	if err != nil {
		t.Fatalf("date.New() unexpected error: %v", err)
	}

	agriculturalYear.SetOpeningDate(openingDate)
	agriculturalYear.SetClosingDate(closingDate)

	gotOpeningDate, ok := agriculturalYear.OpeningDate()
	if !ok {
		t.Fatal("expected opening date to be defined")
	}

	if gotOpeningDate != openingDate {
		t.Errorf(
			"OpeningDate() = %v, want %v",
			gotOpeningDate,
			openingDate,
		)
	}

	gotClosingDate, ok := agriculturalYear.ClosingDate()
	if !ok {
		t.Fatal("expected closing date to be defined")
	}

	if gotClosingDate != closingDate {
		t.Errorf(
			"ClosingDate() = %v, want %v",
			gotClosingDate,
			closingDate,
		)
	}
}
