package user_test

import (
	"errors"
	"os"
	"testing"

	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/errs"
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
	v1, err := NewPassword("sIjjw9WlCa22hVfb")
	assert.Equal(t, [32]uint8{0x5a, 0x9b, 0xab, 0xe4, 0x2e, 0xe8, 0xbb, 0x1b, 0x60, 0x3a, 0x9d, 0xf3, 0x7a, 0x22, 0xef, 0xce, 0xc, 0x63, 0xad, 0x42, 0x70, 0x47, 0x91, 0xea, 0xd3, 0x70, 0x2b, 0x3e, 0xf5, 0x1a, 0x2d, 0x58}, v1.Digest())
	assert.NoError(t, err)
	v2, err := NewPassword("d8D83ffde48bcs74")
	assert.Equal(t, [32]uint8{0x3d, 0x12, 0x6c, 0x41, 0x17, 0xbc, 0x70, 0xe2, 0x50, 0xd6, 0xd1, 0x34, 0xff, 0xe, 0xdb, 0xd3, 0xd, 0xce, 0x61, 0xa5, 0xad, 0x85, 0xf7, 0x60, 0x3f, 0xc4, 0xae, 0xfc, 0x65, 0x58, 0x73, 0x8c}, v2.Digest())
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
	p.Digest()
}
