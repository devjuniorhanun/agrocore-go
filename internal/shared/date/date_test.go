package date

import (
	"errors"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		year    int
		month   time.Month
		day     int
		wantErr error
	}{
		{
			name:  "creates valid date",
			year:  2026,
			month: time.July,
			day:   1,
		},
		{
			name:  "creates leap year date",
			year:  2028,
			month: time.February,
			day:   29,
		},
		{
			name:    "rejects invalid leap year date",
			year:    2026,
			month:   time.February,
			day:     29,
			wantErr: ErrInvalidDate,
		},
		{
			name:    "rejects day outside month",
			year:    2026,
			month:   time.April,
			day:     31,
			wantErr: ErrInvalidDate,
		},
		{
			name:    "rejects invalid month",
			year:    2026,
			month:   time.Month(13),
			day:     1,
			wantErr: ErrInvalidDate,
		},
		{
			name:    "rejects invalid day",
			year:    2026,
			month:   time.January,
			day:     0,
			wantErr: ErrInvalidDate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.year, tt.month, tt.day)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("New() error = %v, want %v", err, tt.wantErr)
			}

			if tt.wantErr != nil {
				return
			}

			if got.Year() != tt.year {
				t.Errorf("Year() = %d, want %d", got.Year(), tt.year)
			}

			if got.Month() != tt.month {
				t.Errorf("Month() = %v, want %v", got.Month(), tt.month)
			}

			if got.Day() != tt.day {
				t.Errorf("Day() = %d, want %d", got.Day(), tt.day)
			}
		})
	}
}

func TestDateComparison(t *testing.T) {
	earlier, err := New(2026, time.July, 1)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	later, err := New(2027, time.June, 30)
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if !earlier.Before(later) {
		t.Error("expected earlier date to be before later date")
	}

	if !later.After(earlier) {
		t.Error("expected later date to be after earlier date")
	}
}
