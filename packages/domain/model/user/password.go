package user

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
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

func (v *Password) Digest(s *PasswordSalt) *[32]byte {
	p := os.Getenv("PEPPER_PASSWORD")
	if len(p) < 1 {
		panic("PEPPER_PASSWORD")
	}
	str := []byte(p + v.value + s.String())
	d := sha256.Sum256(str)
	return &d
}

type PasswordSalt [32]byte

func GenerateSalt() *PasswordSalt {
	r, err := random()
	if err != nil {
		panic(err)
	}
	s := PasswordSalt(*r)
	return &s
}

func (v *PasswordSalt) String() string {
	return string(v[:])
}

func random() (*[32]byte, error) {
	var b [32]byte
	if _, err := rand.Read(b[:]); err != nil {
		return nil, errors.New("unexpected error")
	}
	return &b, nil
}
