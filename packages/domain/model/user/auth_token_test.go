package user_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestAuthToken(t *testing.T) {
	var v *AuthToken
	var err error
	v, err = NewAuthToken("abcdefghijklmnop")
	assert.Equal(t, "abcdefghijklmnop", v.String())
	assert.NoError(t, err)
}

func TestErrorAuthToken(t *testing.T) {
	var v *AuthToken
	var err error
	v, err = NewAuthToken("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
