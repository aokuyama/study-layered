package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"

	"github.com/stretchr/testify/assert"
)

func TestEventID(t *testing.T) {
	var v *CircleID
	var err error
	v, err = NewCircleID("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1")
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", v.String())
	assert.NoError(t, err)
	v, err = NewCircleID("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
