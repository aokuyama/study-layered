package path_test

import (
	"os"
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"1", "sIjjw9WlCa22hVfb"},
		{"2", "d8D83ffde48bcs74"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := NewPath(tt.path)
			assert.Equal(t, "*****", v.String())
			assert.NoError(t, err)
		})
	}
}

func TestErrorPath(t *testing.T) {
	var v *Path
	var err error
	v, err = NewPath("sIjjw9WlCa22hVf")
	assert.Nil(t, v)
	assert.Error(t, err)
	v, err = NewPath("sIjjw9WlCa22hVb12")
	assert.Nil(t, v)
	assert.Error(t, err)
	v, err = NewPath("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}

func TestPathErrorChar(t *testing.T) {
	var v *Path
	var err error
	v, err = NewPath("あいうえおabcde123456")
	assert.Nil(t, v)
	assert.Error(t, err)
	v, err = NewPath("sIjjw9WlCa22hVb/")
	assert.Nil(t, v)
	assert.Error(t, err)
	v, err = NewPath(".sIjjw9WlCa22hVb")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}

func TestGeneratePath(t *testing.T) {
	var err error
	v1, err := GeneratePath()
	assert.Equal(t, "*****", v1.String())
	assert.NoError(t, err)
	v2, err := GeneratePath()
	assert.Equal(t, "*****", v2.String())
	assert.NoError(t, err)
	assert.False(t, v1.Equals(v2))
}

func TestRawValue(t *testing.T) {
	var err error
	v1, err := NewPath("sIjjw9WlCa22hVfb")
	assert.Equal(t, "sIjjw9WlCa22hVfb", v1.RawValue())
	assert.NoError(t, err)
	v2, err := NewPath("d8D83ffde48bcs74")
	assert.Equal(t, "d8D83ffde48bcs74", v2.RawValue())
	assert.NoError(t, err)
	assert.False(t, v1.Equals(v2))
}

func TestDigest(t *testing.T) {
	os.Setenv("PEPPER_PATH", "pepper")

	var err error
	v1, err := NewPath("sIjjw9WlCa22hVfb")
	assert.Equal(t, [32]uint8{0x5a, 0x9b, 0xab, 0xe4, 0x2e, 0xe8, 0xbb, 0x1b, 0x60, 0x3a, 0x9d, 0xf3, 0x7a, 0x22, 0xef, 0xce, 0xc, 0x63, 0xad, 0x42, 0x70, 0x47, 0x91, 0xea, 0xd3, 0x70, 0x2b, 0x3e, 0xf5, 0x1a, 0x2d, 0x58}, v1.Digest())
	assert.NoError(t, err)
	v2, err := NewPath("d8D83ffde48bcs74")
	assert.Equal(t, [32]uint8{0x3d, 0x12, 0x6c, 0x41, 0x17, 0xbc, 0x70, 0xe2, 0x50, 0xd6, 0xd1, 0x34, 0xff, 0xe, 0xdb, 0xd3, 0xd, 0xce, 0x61, 0xa5, 0xad, 0x85, 0xf7, 0x60, 0x3f, 0xc4, 0xae, 0xfc, 0x65, 0x58, 0x73, 0x8c}, v2.Digest())
	assert.NoError(t, err)
	assert.False(t, v1.Equals(v2))
}

func TestPanicDigest(t *testing.T) {
	os.Setenv("PEPPER_PATH", "")

	p, _ := NewPath("sIjjw9WlCa22hVfb")

	defer func() {
		err := recover()
		if err != "PEPPER_PATH" {
			t.Errorf("got %v\nwant %v", err, "PEPPER_PATH")
		}
	}()
	p.Digest()
}

func TestEncrypt(t *testing.T) {
	os.Setenv("KEY_PATH", "1234567890abcdef")

	tests := []struct {
		name string
		path string
	}{
		{"1", "sIjjw9WlCa22hVfb"},
		{"2", "d8D83ffde48bcs74"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p1, _ := NewPath(tt.path)
			e1, err := p1.Encrypt()
			assert.NoError(t, err)
			p1_2, err := DecryptPath(e1)
			assert.NoError(t, err)
			assert.Equal(t, tt.path, p1_2.RawValue())
		})
	}
}

func TestErrorEncrypt(t *testing.T) {
	os.Setenv("KEY_PATH", "")

	p, _ := NewPath("sIjjw9WlCa22hVfb")

	e1, err := p.Encrypt()
	assert.Nil(t, e1)
	assert.Error(t, err)
	assert.Equal(t, "KEY_PATH", err.Error())

	e2, err := DecryptPath(&Encrypted{})
	assert.Nil(t, e2)
	assert.Error(t, err)
	assert.Equal(t, "KEY_PATH", err.Error())
}
