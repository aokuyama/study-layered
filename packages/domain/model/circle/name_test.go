package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

	"github.com/stretchr/testify/assert"
)

func TestName(t *testing.T) {
	var v *Name
	var err error
	v, err = NewName("abcdefghijklmnop")
	assert.Equal(t, "abcdefghijklmnop", v.String())
	assert.NoError(t, err)
	v, err = NewName("あいうえお")
	assert.Equal(t, "あいうえお", v.String())
	assert.NoError(t, err)
}

func TestErrorName(t *testing.T) {
	var v *Name
	var err error
	v, err = NewName("abcdefghijklmnopq")
	assert.Nil(t, v)
	assert.Error(t, err, "over 16 characters")
	v, err = NewName("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
