// Package main demonstrates advanced error handling: wrapping errors with
// context, and inspecting error chains with errors.Is and errors.As.
package main

import (
	"errors"
	"fmt"
)

// ErrNotFound is a sentinel error other code can check for with errors.Is,
// no matter how many layers of context wrap it along the way.
var ErrNotFound = errors.New("resource not found")

// ValidationError is a custom error type carrying structured data, meant
// to be recovered from an error chain with errors.As.
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field %q", e.Field)
}

func lookupUser(id int) error {
	if id < 0 {
		return &ValidationError{Field: "id"}
	}
	if id != 1 {
		// %w wraps ErrNotFound, preserving it in the chain instead of
		// discarding it behind a plain string.
		return fmt.Errorf("lookupUser(%d): %w", id, ErrNotFound)
	}
	return nil
}

func loadProfile(id int) error {
	if err := lookupUser(id); err != nil {
		// Each layer adds its own context without destroying the root cause.
		return fmt.Errorf("loadProfile: %w", err)
	}
	return nil
}

func main() {
	err := loadProfile(42)
	fmt.Println("error:", err)
	fmt.Println("is ErrNotFound:", errors.Is(err, ErrNotFound))

	err = loadProfile(-1)
	fmt.Println("error:", err)

	var valErr *ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("recovered ValidationError for field:", valErr.Field)
	}

	if err := loadProfile(1); err == nil {
		fmt.Println("loadProfile(1) succeeded, no error")
	}
}
