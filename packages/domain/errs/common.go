package errs

import (
	"errors"
	"fmt"
)

var ErrBadParam = errors.New("bad parametaer")
var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
var ErrConflict = errors.New("conflict")
var ErrFatal = errors.New("fatal error")

func NewBadParam(text string) error {
	return fmt.Errorf("%w: %s", ErrBadParam, text)
}

func NewNotFound(text string) error {
	return fmt.Errorf("%w: %s", ErrNotFound, text)
}

func NewUnauthorized(text string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, text)
}

func NewConflict(text string) error {
	return fmt.Errorf("%w: %s", ErrConflict, text)
}

func NewFatal(text string) error {
	return fmt.Errorf("%w: %s", ErrFatal, text)
}
