package path

import (
	"crypto/rand"
	"crypto/sha256"
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

func GeneratePath() (*Path, error) {
	str, err := randomStr(16)
	if err != nil {
		return nil, err
	}
	return NewPath(str)
}

func (v *Path) String() string {
	return "*****"
}

func (v *Path) RawValue() string {
	return v.value
}

func (v *Path) Digest() [32]byte {
	return sha256.Sum256([]byte(v.value))
}

func (v *Path) Equals(p *Path) bool {
	return v.value == p.value
}

func randomStr(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
