package owner_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/owner"

	"github.com/stretchr/testify/assert"
)

func TestOwnerID(t *testing.T) {
	var v *OwnerID
	var err error
	v, err = NewOwnerID("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1")
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", v.String())
	assert.NoError(t, err)
	v, err = NewOwnerID("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
