package circle_test

import (
	"testing"

	. "github.com/aokuyama/circle_scheduler-api/packages/domain/model/circle"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/model/common/path"
	"github.com/aokuyama/circle_scheduler-api/packages/domain/test"

	"github.com/stretchr/testify/assert"
)

func TestNewEntity(t *testing.T) {
	var e *Circle
	var err error
	i := "26f90f21-dd19-4df1-81ff-ea9dcbcf03d1"
	o := "d833a112-95e8-4042-ab02-ffde48bc874a"
	n := "circle"
	p := path.Path{}

	e, err = NewCircle(&i, &o, &n, &p)
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
			p := path.Path{}
			e, err := NewCircle(&tt.id, &tt.ownerID, &tt.name, &p)
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
	p := path.Path{}
	e1 := test.PanicOr(NewCircle(&i1, &oi, &n, &p))
	e2 := test.PanicOr(NewCircle(&i1, &oi, &n, &p))
	e3 := test.PanicOr(NewCircle(&i2, &oi, &n, &p))
	assert.True(t, e1.Identical(e2))
	assert.False(t, e1.Identical(e3))
}
