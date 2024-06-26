package path

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"os"
	"regexp"
	"unicode/utf8"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
)

type Path struct {
	value string
}

var reg = regexp.MustCompile("^[0-9A-Za-z]{16}$")

func NewPath(v string) (*Path, error) {
	c := utf8.RuneCountInString(v)
	if c != 16 {
		return nil, errs.NewBadParam("must 16 characters")
	}
	if !reg.MatchString(v) {
		return nil, errs.NewBadParam("alphanumeric only")
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
	p := os.Getenv("PEPPER_PATH")
	if len(p) < 1 {
		panic("PEPPER_PATH")
	}
	s := []byte(p + v.value)
	return sha256.Sum256(s)
}

func (v *Path) Encrypt() (*Encrypted, error) {
	k := os.Getenv("KEY_PATH")
	if len(k) < 1 {
		return nil, errors.New("KEY_PATH")
	}
	return encrypt(v.value, k)
}

func DecryptPath(e *Encrypted) (*Path, error) {
	k := os.Getenv("KEY_PATH")
	if len(k) < 1 {
		return nil, errors.New("KEY_PATH")
	}

	p, err := decrypt(e, k)

	if err != nil {
		return nil, err
	}
	return NewPath(*p)
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
