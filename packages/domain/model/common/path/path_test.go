package path_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	var v *Path
	var err error
	v, err = NewPath("sIjjw9WlCa22hVfb")
	assert.Equal(t, "*****", v.String())
	assert.NoError(t, err)
	v, err = NewPath("d8D83ffde48bcs74")
	assert.Equal(t, "*****", v.String())
	assert.NoError(t, err)
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

func TestDigest(t *testing.T) {
	var err error
	v1, err := NewPath("sIjjw9WlCa22hVfb")
	assert.Equal(t, [32]uint8{0xf, 0xeb, 0x5a, 0x1f, 0xee, 0x86, 0x33, 0x6, 0x32, 0x84, 0x5e, 0x9d, 0xf7, 0x34, 0x14, 0xb0, 0xad, 0x56, 0x73, 0x85, 0x3, 0x9e, 0x72, 0x5a, 0x49, 0xf8, 0x97, 0xb9, 0xb9, 0xf5, 0xb, 0xbf}, v1.Digest())
	assert.NoError(t, err)
	v2, err := NewPath("d8D83ffde48bcs74")
	assert.Equal(t, [32]uint8{0x9b, 0x6e, 0xea, 0xfb, 0x6e, 0x6d, 0x5c, 0x5c, 0xe6, 0xb7, 0x18, 0x1b, 0xea, 0xf8, 0xfa, 0xbd, 0xe3, 0x7a, 0xff, 0x77, 0x78, 0x12, 0x1d, 0x5f, 0xcb, 0x20, 0x6a, 0x21, 0xfa, 0x3f, 0x2e, 0xa4}, v2.Digest())
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
