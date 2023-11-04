package user_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/user"

	"github.com/stretchr/testify/assert"
)

func TestUserID(t *testing.T) {
	var v *UserID
	var err error
	v, err = NewUserID("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1")
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", v.String())
	assert.NoError(t, err)
	v, err = NewUserID("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
