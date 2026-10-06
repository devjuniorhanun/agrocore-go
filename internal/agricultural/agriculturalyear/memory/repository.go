package memory

import (
	"context"
	"sync"

	"github.com/devjuniorhanun/agrocore-go/internal/agricultural/agriculturalyear"
)

// Repository stores agricultural years in memory.
//
// It is intended for development, tests, and scenarios that do not require
// durable persistence.
type Repository struct {
	mu     sync.RWMutex
	nextID int64
	items  map[int64]agriculturalyear.AgriculturalYear
}

// NewRepository creates an empty in-memory agricultural year repository.
func NewRepository() *Repository {
	return &Repository{
		nextID: 1,
		items:  make(map[int64]agriculturalyear.AgriculturalYear),
	}
}

// ExistsByName reports whether an agricultural year with the given name exists.
func (r *Repository) ExistsByName(
	ctx context.Context,
	name string,
) (bool, error) {
	if err := ctx.Err(); err != nil {
		return false, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, item := range r.items {
		if item.Name() == name {
			return true, nil
		}
	}

	return false, nil
}

// Save stores an agricultural year and assigns a persistence identifier.
func (r *Repository) Save(
	ctx context.Context,
	value agriculturalyear.AgriculturalYear,
) (agriculturalyear.AgriculturalYear, error) {
	if err := ctx.Err(); err != nil {
		return agriculturalyear.AgriculturalYear{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	saved, err := agriculturalyear.Restore(
		r.nextID,
		value.Name(),
		value.Status(),
	)
	if err != nil {
		return agriculturalyear.AgriculturalYear{}, err
	}

	if openingDate, ok := value.OpeningDate(); ok {
		saved.SetOpeningDate(openingDate)
	}

	if closingDate, ok := value.ClosingDate(); ok {
		saved.SetClosingDate(closingDate)
	}

	r.items[r.nextID] = saved
	r.nextID++

	return saved, nil
}
