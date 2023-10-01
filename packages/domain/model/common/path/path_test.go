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
