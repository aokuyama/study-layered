package user

import (
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
	"unicode/utf8"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
)

type Password struct {
	value string
}

var reg = regexp.MustCompile("^[a-zA-Z0-9!-/:-@¥[-`{-~]*$")

func NewPassword(v string) (*Password, error) {
	c := utf8.RuneCountInString(v)
	if c < 12 {
		return nil, fmt.Errorf("%w can`t be less than 12 chars", errs.ErrBadParam)
	}
	if !reg.MatchString(v) {
		return nil, fmt.Errorf("%w half-with characters only", errs.ErrBadParam)
	}
	p := Password{v}
	return &p, nil
}

func (v *Password) String() string {
	return "*****"
}

func (v *Password) Digest() [32]byte {
	p := os.Getenv("PEPPER_PASSWORD")
	if len(p) < 1 {
		panic("PEPPER_PASSWORD")
	}
	s := []byte(p + v.value)
	return sha256.Sum256(s)
}
