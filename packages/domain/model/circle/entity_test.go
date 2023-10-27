package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/util"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	var e *CircleEntity
	var err error
	i := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	o := "d833a112-95e8-4042-ab02-ffde48bc874a"
	n := "circle"

	e, err = NewCircleEntity(&i, &o, &n)
	assert.Equal(t, "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", e.ID().String())
	assert.Equal(t, "d833a112-95e8-4042-ab02-ffde48bc874a", e.OwnerID().String())
	assert.Equal(t, "circle", e.Name().String())
	assert.NoError(t, err)
}

func TestErrorNewEntity(t *testing.T) {
	tests := []struct {
		testName          string
		id, ownerID, name string
	}{
		{"id", "invalid", "d833a112-95e8-4042-ab02-ffde48bc874a", "circle"},
		{"ownerID", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "invalid", "circle"},
		{"name", "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1", "d833a112-95e8-4042-ab02-ffde48bc874a", ""},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			e, err := NewCircleEntity(&tt.id, &tt.ownerID, &tt.name)
			assert.Nil(t, e)
			assert.Error(t, err)
		})
	}
}

func TestIdenticalEntity(t *testing.T) {
	n := "a"
	i1 := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	i2 := "550e8400-e29b-41d4-a716-446655440000"
	oi := "d833a112-95e8-4042-ab02-ffde48bc874a"
	e1 := util.PanicOr(NewCircleEntity(&i1, &oi, &n))
	e2 := util.PanicOr(NewCircleEntity(&i1, &oi, &n))
	e3 := util.PanicOr(NewCircleEntity(&i2, &oi, &n))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
