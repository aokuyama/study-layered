package path

import (
	"errors"
	"regexp"
	"unicode/utf8"
)

type Path struct {
	value string
}

var reg = regexp.MustCompile("[0-9A-Za-z]{16}$")

func NewPath(v string) (*Path, error) {
	c := utf8.RuneCountInString(v)
	if c != 16 {
		return nil, errors.New("must 16 characters")
	}
	if !reg.MatchString(v) {
		return nil, errors.New("alphanumeric only")
	}
	p := Path{v}
	return &p, nil
}

func (v *Path) String() string {
	return "*****"
}
