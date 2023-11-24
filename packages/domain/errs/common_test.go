package errs_test

import (
	"errors"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	"github.com/go-playground/assert/v2"
)

func TestError(t *testing.T) {
	tests := []struct {
		name                                              string
		newErr                                            func(string) error
		badParam, notFound, unauthorized, conflict, fatal bool
	}{
		{"badParam", NewBadParam, true, false, false, false, false},
		{"notFound", NewNotFound, false, true, false, false, false},
		{"unauthorized", NewUnauthorized, false, false, true, false, false},
		{"conflict", NewConflict, false, false, false, true, false},
		{"fatal", NewFatal, false, false, false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.newErr("")
			assert.Equal(t, tt.badParam, errors.Is(err, ErrBadParam))
			assert.Equal(t, tt.notFound, errors.Is(err, ErrNotFound))
			assert.Equal(t, tt.unauthorized, errors.Is(err, ErrUnauthorized))
			assert.Equal(t, tt.conflict, errors.Is(err, ErrConflict))
			assert.Equal(t, tt.fatal, errors.Is(err, ErrFatal))
		})
	}
}
