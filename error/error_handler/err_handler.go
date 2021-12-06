package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	ErrJWTMissing = errors.New("jwt missing")
	ErrJWTInvalid = errors.New("jwt invalid")
)

type ErrorHandler func(w http.ResponseWriter, r *http.Request, err error)

type invalidError struct {
	details error
}

func (e *invalidError) Is(target error) bool {
	return target == ErrJWTInvalid
}

func (e *invalidError) Error() string {
	return fmt.Sprintf("%s: %s", ErrJWTInvalid, e.details)
}

func (e *invalidError) Unwrap() error {
	return e.details
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	switch {
	case errors.Is(err, ErrJWTMissing):
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"message":"JWT is missing."}`))
	case errors.Is(err, ErrJWTInvalid):
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"message":"JWT is invalid."}`))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"Internal error checking the JWT."}`))
	}
}
