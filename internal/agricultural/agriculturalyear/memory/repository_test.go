package memory

import (
	"context"
	"errors"
	"testing"

	"github.com/devjuniorhanun/agrocore-go/internal/agricultural/agriculturalyear"
)

func TestRepositorySave(t *testing.T) {
	repository := NewRepository()

	value, err := agriculturalyear.New("2026/2027")
	if err != nil {
		t.Fatalf("agriculturalyear.New() unexpected error: %v", err)
	}

	saved, err := repository.Save(
		context.Background(),
		value,
	)
	if err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	if saved.ID() != 1 {
		t.Errorf("ID() = %d, want 1", saved.ID())
	}

	if saved.Name() != "2026/2027" {
		t.Errorf(
			"Name() = %q, want %q",
			saved.Name(),
			"2026/2027",
		)
	}
}

func TestRepositorySaveIncrementsID(t *testing.T) {
	repository := NewRepository()

	first, err := agriculturalyear.New("2025/2026")
	if err != nil {
		t.Fatalf("agriculturalyear.New() unexpected error: %v", err)
	}

	second, err := agriculturalyear.New("2026/2027")
	if err != nil {
		t.Fatalf("agriculturalyear.New() unexpected error: %v", err)
	}

	firstSaved, err := repository.Save(
		context.Background(),
		first,
	)
	if err != nil {
		t.Fatalf("first Save() unexpected error: %v", err)
	}

	secondSaved, err := repository.Save(
		context.Background(),
		second,
	)
	if err != nil {
		t.Fatalf("second Save() unexpected error: %v", err)
	}

	if firstSaved.ID() != 1 {
		t.Errorf(
			"first ID() = %d, want 1",
			firstSaved.ID(),
		)
	}

	if secondSaved.ID() != 2 {
		t.Errorf(
			"second ID() = %d, want 2",
			secondSaved.ID(),
		)
	}
}

func TestRepositoryExistsByName(t *testing.T) {
	repository := NewRepository()

	value, err := agriculturalyear.New("2026/2027")
	if err != nil {
		t.Fatalf("agriculturalyear.New() unexpected error: %v", err)
	}

	_, err = repository.Save(
		context.Background(),
		value,
	)
	if err != nil {
		t.Fatalf("Save() unexpected error: %v", err)
	}

	exists, err := repository.ExistsByName(
		context.Background(),
		"2026/2027",
	)
	if err != nil {
		t.Fatalf("ExistsByName() unexpected error: %v", err)
	}

	if !exists {
		t.Error("ExistsByName() = false, want true")
	}

	exists, err = repository.ExistsByName(
		context.Background(),
		"2027/2028",
	)
	if err != nil {
		t.Fatalf("ExistsByName() unexpected error: %v", err)
	}

	if exists {
		t.Error("ExistsByName() = true, want false")
	}
}

func TestRepositoryRespectsCanceledContext(t *testing.T) {
	repository := NewRepository()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := repository.ExistsByName(
		ctx,
		"2026/2027",
	)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf(
			"ExistsByName() error = %v, want %v",
			err,
			context.Canceled,
		)
	}

	value, err := agriculturalyear.New("2026/2027")
	if err != nil {
		t.Fatalf("agriculturalyear.New() unexpected error: %v", err)
	}

	_, err = repository.Save(ctx, value)

	if !errors.Is(err, context.Canceled) {
		t.Fatalf(
			"Save() error = %v, want %v",
			err,
			context.Canceled,
		)
	}
}
