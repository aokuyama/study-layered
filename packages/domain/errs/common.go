package errs

import "errors"

var ErrBadParam = errors.New("bad parametaer")
var ErrNotFound = errors.New("not found")
var ErrFatal = errors.New("fatal error")
var ErrUnauthorized = errors.New("unauthorized")
var ErrConflict = errors.New("conflict")
