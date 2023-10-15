package event

import (
	"errors"
	"unicode/utf8"
)

type Name string

func NewName(v string) (*Name, error) {
	c := utf8.RuneCountInString(v)
	if c == 0 {
		return nil, errors.New("can`t be blank")
	}
	if c > 16 {
		return nil, errors.New("enter within 16 characters")
	}
	i := Name(v)
	return &i, nil
}

func (v *Name) String() string {
	return string(*v)
}

func (v *Name) Equals(n *Name) bool {
	return v.String() == n.String()
}
