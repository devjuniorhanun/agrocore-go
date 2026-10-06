package agriculturalyear

import (
	"errors"
	"strings"

	"github.com/devjuniorhanun/agrocore-go/internal/shared/date"
)

// Status represents the status of an agricultural year.
type Status string

const (
	// StatusActive represents an active agricultural year.
	StatusActive Status = "A"
)

// ErrNameRequired is returned when the agricultural year name is empty.
var ErrNameRequired = errors.New("agricultural year name is required")

// AgriculturalYear represents an agricultural year in the domain.
type AgriculturalYear struct {
	id          int64
	name        string
	openingDate *date.Date
	closingDate *date.Date
	status      Status
}

// New creates a new active agricultural year without a persistence identifier.
func New(name string) (AgriculturalYear, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return AgriculturalYear{}, ErrNameRequired
	}

	return AgriculturalYear{
		name:   name,
		status: StatusActive,
	}, nil
}

// Restore recreates an agricultural year that already has a persistence
// identifier.
//
// It is intended for infrastructure adapters that need to reconstruct an
// entity after persistence or when reading stored data.
func Restore(id int64, name string, status Status) (AgriculturalYear, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return AgriculturalYear{}, ErrNameRequired
	}

	return AgriculturalYear{
		id:     id,
		name:   name,
		status: status,
	}, nil
}

// ID returns the agricultural year identifier.
//
// A zero value means that the agricultural year has not yet received
// a persistence identifier.
func (a AgriculturalYear) ID() int64 {
	return a.id
}

// Name returns the agricultural year name.
func (a AgriculturalYear) Name() string {
	return a.name
}

// Status returns the agricultural year status.
func (a AgriculturalYear) Status() Status {
	return a.status
}

// SetOpeningDate defines the agricultural year opening date.
func (a *AgriculturalYear) SetOpeningDate(value date.Date) {
	a.openingDate = &value
}

// SetClosingDate defines the agricultural year closing date.
func (a *AgriculturalYear) SetClosingDate(value date.Date) {
	a.closingDate = &value
}

// OpeningDate returns the agricultural year opening date when defined.
func (a AgriculturalYear) OpeningDate() (date.Date, bool) {
	if a.openingDate == nil {
		return date.Date{}, false
	}

	return *a.openingDate, true
}

// ClosingDate returns the agricultural year closing date when defined.
func (a AgriculturalYear) ClosingDate() (date.Date, bool) {
	if a.closingDate == nil {
		return date.Date{}, false
	}

	return *a.closingDate, true
}
