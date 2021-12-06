package main

import (
	"github.com/pkg/errors"
	"testing"
)

func Test_invalidError(t *testing.T) {
	t.Run("Is", func(t *testing.T) {
		e := invalidError{details: errors.New("error details")}
		if !errors.Is(&e, ErrJWTInvalid) {
			t.Fatal("expected invalidError is ErrJWTInvalid, but is was not")
		}
	})

	t.Run("Error", func(t *testing.T) {
		e := invalidError{details: errors.New("error details")}
		assertErrorMsg(t, "jwt invalid: error details", &e)
	})

	t.Run("Unwrap", func(t *testing.T) {
		expectedErr := errors.New("expected err")
		e := invalidError{details: expectedErr}

		if !errors.Is(&e, expectedErr) {
			t.Fatal("expected invalidError is expectedErr, but it was not")
		}
	})
}

func assertErrorMsg(t testing.TB, expected string, got error) {
	if (expected == "" && got != nil) || (expected != "" && (got == nil || got.Error() != expected)) {
		t.Fatalf("expected error: %s, got %v", expected, got)
	}
}
