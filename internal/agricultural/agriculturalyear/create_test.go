package agriculturalyear

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/devjuniorhanun/agrocore-go/internal/shared/date"
)

type fakeRepository struct {
	exists      bool
	existsErr   error
	saveErr     error
	checkedName string
	saved       AgriculturalYear
	nextID      int64
}

func (r *fakeRepository) ExistsByName(
	_ context.Context,
	name string,
) (bool, error) {
	r.checkedName = name

	if r.existsErr != nil {
		return false, r.existsErr
	}

	return r.exists, nil
}

func (r *fakeRepository) Save(
	_ context.Context,
	agriculturalYear AgriculturalYear,
) (AgriculturalYear, error) {
	if r.saveErr != nil {
		return AgriculturalYear{}, r.saveErr
	}

	if r.nextID == 0 {
		r.nextID = 1
	}

	saved, err := Restore(
		r.nextID,
		agriculturalYear.Name(),
		agriculturalYear.Status(),
	)
	if err != nil {
		return AgriculturalYear{}, err
	}

	if openingDate, ok := agriculturalYear.OpeningDate(); ok {
		saved.SetOpeningDate(openingDate)
	}

	if closingDate, ok := agriculturalYear.ClosingDate(); ok {
		saved.SetClosingDate(closingDate)
	}

	r.saved = saved

	return r.saved, nil
}

func TestCreate(t *testing.T) {
	repository := &fakeRepository{
		nextID: 10,
	}

	create := NewCreator(repository)

	got, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "2026/2027",
		},
	)

	if err != nil {
		t.Fatalf("Execute() unexpected error: %v", err)
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

	if repository.saved.ID() != 10 {
		t.Errorf(
			"saved ID = %d, want 10",
			repository.saved.ID(),
		)
	}
}

func TestCreateWithDates(t *testing.T) {
	repository := &fakeRepository{
		nextID: 10,
	}

	openingDate, err := date.New(
		2026,
		time.July,
		1,
	)
	if err != nil {
		t.Fatalf(
			"date.New() unexpected error: %v",
			err,
		)
	}

	closingDate, err := date.New(
		2027,
		time.June,
		30,
	)
	if err != nil {
		t.Fatalf(
			"date.New() unexpected error: %v",
			err,
		)
	}

	create := NewCreator(repository)

	got, err := create.Execute(
		context.Background(),
		CreateInput{
			Name:        "2026/2027",
			OpeningDate: &openingDate,
			ClosingDate: &closingDate,
		},
	)
	if err != nil {
		t.Fatalf("Execute() unexpected error: %v", err)
	}

	gotOpeningDate, ok := got.OpeningDate()
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

	gotClosingDate, ok := got.ClosingDate()
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

func TestCreateWithoutDates(t *testing.T) {
	repository := &fakeRepository{}

	create := NewCreator(repository)

	got, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "2026/2027",
		},
	)
	if err != nil {
		t.Fatalf("Execute() unexpected error: %v", err)
	}

	if _, ok := got.OpeningDate(); ok {
		t.Error("expected opening date to be undefined")
	}

	if _, ok := got.ClosingDate(); ok {
		t.Error("expected closing date to be undefined")
	}
}

func TestCreateRejectsDuplicateName(t *testing.T) {
	repository := &fakeRepository{
		exists: true,
	}

	create := NewCreator(repository)

	_, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "2026/2027",
		},
	)

	if !errors.Is(err, ErrNameAlreadyExists) {
		t.Fatalf(
			"Execute() error = %v, want %v",
			err,
			ErrNameAlreadyExists,
		)
	}
}

func TestCreateChecksNormalizedName(t *testing.T) {
	repository := &fakeRepository{}

	create := NewCreator(repository)

	_, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "  2026/2027  ",
		},
	)

	if err != nil {
		t.Fatalf("Execute() unexpected error: %v", err)
	}

	if repository.checkedName != "2026/2027" {
		t.Errorf(
			"checked name = %q, want %q",
			repository.checkedName,
			"2026/2027",
		)
	}
}

func TestCreateReturnsRepositoryExistsError(t *testing.T) {
	repositoryError := errors.New("repository error")

	repository := &fakeRepository{
		existsErr: repositoryError,
	}

	create := NewCreator(repository)

	_, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "2026/2027",
		},
	)

	if !errors.Is(err, repositoryError) {
		t.Fatalf(
			"Execute() error = %v, want %v",
			err,
			repositoryError,
		)
	}
}

func TestCreateReturnsRepositorySaveError(t *testing.T) {
	repositoryError := errors.New("repository error")

	repository := &fakeRepository{
		saveErr: repositoryError,
	}

	create := NewCreator(repository)

	_, err := create.Execute(
		context.Background(),
		CreateInput{
			Name: "2026/2027",
		},
	)

	if !errors.Is(err, repositoryError) {
		t.Fatalf(
			"Execute() error = %v, want %v",
			err,
			repositoryError,
		)
	}
}
