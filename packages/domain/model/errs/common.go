package errs

import "errors"

var ErrBadParam = errors.New("bad parametaer")
var ErrNotFound = errors.New("not found")
var ErrFatal = errors.New("fatal error")
