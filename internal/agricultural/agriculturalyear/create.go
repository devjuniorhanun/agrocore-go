package agriculturalyear

import (
	"context"
	"errors"

	"github.com/devjuniorhanun/agrocore-go/internal/shared/date"
)

// ErrNameAlreadyExists is returned when an agricultural year with the same
// normalized name already exists.
var ErrNameAlreadyExists = errors.New(
	"agricultural year name already exists",
)

// Repository defines the persistence capabilities required by agricultural
// year application operations.
type Repository interface {
	ExistsByName(ctx context.Context, name string) (bool, error)
	Save(
		ctx context.Context,
		agriculturalYear AgriculturalYear,
	) (AgriculturalYear, error)
}

// CreateInput contains the data required to create an agricultural year.
type CreateInput struct {
	Name        string
	OpeningDate *date.Date
	ClosingDate *date.Date
}

// Creator creates agricultural years.
type Creator struct {
	repository Repository
}

// NewCreator creates an agricultural year Creator.
func NewCreator(repository Repository) *Creator {
	return &Creator{
		repository: repository,
	}
}

// Execute creates and persists an agricultural year.
func (c *Creator) Execute(
	ctx context.Context,
	input CreateInput,
) (AgriculturalYear, error) {
	agriculturalYear, err := New(input.Name)
	if err != nil {
		return AgriculturalYear{}, err
	}

	if input.OpeningDate != nil {
		agriculturalYear.SetOpeningDate(*input.OpeningDate)
	}

	if input.ClosingDate != nil {
		agriculturalYear.SetClosingDate(*input.ClosingDate)
	}

	exists, err := c.repository.ExistsByName(
		ctx,
		agriculturalYear.Name(),
	)
	if err != nil {
		return AgriculturalYear{}, err
	}

	if exists {
		return AgriculturalYear{}, ErrNameAlreadyExists
	}

	saved, err := c.repository.Save(ctx, agriculturalYear)
	if err != nil {
		return AgriculturalYear{}, err
	}

	return saved, nil
}
