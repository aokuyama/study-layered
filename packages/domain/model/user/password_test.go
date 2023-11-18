package user_test

import (
	"errors"
	"os"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/errs"
	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"
	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"1", "sIjjw9WlCa22hVfb"},
		{"2", "d8D83ffde48bcs74"},
		{"3", "!-/:-@¥[-`{-~]*"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := NewPassword(tt.path)
			assert.Equal(t, "*****", v.String())
			assert.NoError(t, err)
		})
	}
}

func TestErrorPassword(t *testing.T) {
	var v *Password
	var err error
	v, err = NewPassword("...........")
	assert.Nil(t, v)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, errs.ErrBadParam))
	v, err = NewPassword("")
	assert.Nil(t, v)
	assert.Error(t, err)
	assert.True(t, errors.Is(err, errs.ErrBadParam))
}

func TestPasswordErrorChar(t *testing.T) {
	var v *Password
	var err error
	v, err = NewPassword("あいうえおabcde123456")
	assert.Nil(t, v)
	assert.Error(t, err)
}

func TestDigest(t *testing.T) {
	os.Setenv("PEPPER_PASSWORD", "pepper")

	var err error
	s := PasswordSalt{}
	v1, err := NewPassword("sIjjw9WlCa22hVfb")
	assert.Equal(t, [32]uint8{0xd2, 0xfe, 0x5c, 0xb, 0x5a, 0x3f, 0x30, 0x7f, 0x83, 0xe4, 0x2, 0xb4, 0x4d, 0xae, 0x7b, 0x6a, 0x13, 0xe6, 0x60, 0xde, 0xdd, 0xfa, 0xc5, 0x91, 0x33, 0x7d, 0xbe, 0x57, 0x53, 0x1b, 0x10, 0xc2}, *v1.Digest(&s))
	assert.NoError(t, err)
	v2, err := NewPassword("d8D83ffde48bcs74")
	assert.Equal(t, [32]uint8{0x5f, 0x55, 0x2b, 0xc4, 0xb2, 0xcf, 0xf3, 0x53, 0xd1, 0xf9, 0x4e, 0xbc, 0xc8, 0x3c, 0x20, 0x9f, 0x57, 0x45, 0xd3, 0x1c, 0x1c, 0xa2, 0x73, 0x75, 0xbf, 0x28, 0xb4, 0x30, 0x30, 0x83, 0xb9, 0x26}, *v2.Digest(&s))
	assert.NoError(t, err)
}

func TestPanicDigest(t *testing.T) {
	os.Setenv("PEPPER_PASSWORD", "")

	p, _ := NewPassword("sIjjw9WlCa22hVfb")

	defer func() {
		err := recover()
		if err != "PEPPER_PASSWORD" {
			t.Errorf("got %v\nwant %v", err, "PEPPER_PASSWORD")
		}
	}()
	p.Digest(GenerateSalt())
}

func TestSalt(t *testing.T) {
	s := GenerateSalt()
	assert.Equal(t, 32, len(s.String()))
}
