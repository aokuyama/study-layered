package common_test

import (
	"testing"

	. "aokuyama/circle_scheduler-api/packages/domain/model/common"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	var v *UUID
	var err error
	v, err = NewUUID("26f90f21-dd19-4df1-81ff-ea9dcbcf03d1")
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", v.String())
	assert.NoError(t, err)
	v, err = NewUUID("d833a112-95e8-4042-ab02-ffde48bc874a")
	assert.Equal(t, "d833a112-95e8-4042-ab02-ffde48bc874a", v.String())
	assert.NoError(t, err)
}

func TestErrorUUID(t *testing.T) {
	var v *UUID
	var err error
	v, err = NewUUID("abc")
	assert.Nil(t, v)
	assert.Error(t, err, "deny not uuid string")
	v, err = NewUUID("")
	assert.Nil(t, v)
	assert.Error(t, err, "deny empty")
}
