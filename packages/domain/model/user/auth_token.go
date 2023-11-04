package user

import (
	"errors"
	"unicode/utf8"
)

type AuthToken string

func NewAuthToken(v string) (*AuthToken, error) {
	c := utf8.RuneCountInString(v)
	if c == 0 {
		return nil, errors.New("can`t be blank")
	}
	t := AuthToken(v)
	return &t, nil
}

func (v *AuthToken) String() string {
	return string(*v)
}
