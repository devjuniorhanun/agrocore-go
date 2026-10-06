package agriculturalyear_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/devjuniorhanun/agrocore-go/internal/agricultural/agriculturalyear"
	"github.com/devjuniorhanun/agrocore-go/internal/agricultural/agriculturalyear/memory"
	"github.com/devjuniorhanun/agrocore-go/internal/shared/date"
)

func TestCreateAgriculturalYearWithMemoryRepository(t *testing.T) {
	repository := memory.NewRepository()
	creator := agriculturalyear.NewCreator(repository)
	ctx := context.Background()

	first, err := creator.Execute(
		ctx,
		agriculturalyear.CreateInput{
			Name: "2025/2026",
		},
	)
	if err != nil {
		t.Fatalf("first Execute() unexpected error: %v", err)
	}

	if first.ID() != 1 {
		t.Errorf(
			"first ID() = %d, want 1",
			first.ID(),
		)
	}

	if first.Name() != "2025/2026" {
		t.Errorf(
			"first Name() = %q, want %q",
			first.Name(),
			"2025/2026",
		)
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

	second, err := creator.Execute(
		ctx,
		agriculturalyear.CreateInput{
			Name:        "2026/2027",
			OpeningDate: &openingDate,
			ClosingDate: &closingDate,
		},
	)
	if err != nil {
		t.Fatalf("second Execute() unexpected error: %v", err)
	}

	if second.ID() != 2 {
		t.Errorf(
			"second ID() = %d, want 2",
			second.ID(),
		)
	}

	if second.Name() != "2026/2027" {
		t.Errorf(
			"second Name() = %q, want %q",
			second.Name(),
			"2026/2027",
		)
	}

	gotOpeningDate, ok := second.OpeningDate()
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

	gotClosingDate, ok := second.ClosingDate()
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

	_, err = creator.Execute(
		ctx,
		agriculturalyear.CreateInput{
			Name: "2026/2027",
		},
	)

	if !errors.Is(err, agriculturalyear.ErrNameAlreadyExists) {
		t.Fatalf(
			"duplicate Execute() error = %v, want %v",
			err,
			agriculturalyear.ErrNameAlreadyExists,
		)
	}
}
